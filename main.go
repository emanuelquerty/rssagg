package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/emanuelquerty/rssagg/internal/database"
	"github.com/emanuelquerty/rssagg/internal/handlers"
	"github.com/emanuelquerty/rssagg/internal/routers"
	"github.com/emanuelquerty/rssagg/internal/services"
)

func main() {
	logger := log.New(os.Stdout, "rssagg-api", log.LstdFlags)
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		logger.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		logger.Fatal("DB_URL is not found in the environment")
	}

	dbConn := database.Config(dbURL)

	// Long Running job
	serviceContext := services.NewServiceContext(*logger, dbConn)
	go serviceContext.StartScraping(10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	handlerCtx := handlers.NewHandlerContext(dbConn, logger)
	routerCtx := routers.NewRouterContext(handlerCtx)
	router.Mount("/v1", routerCtx.Route())

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go func() {
		logger.Printf("Server listening on port %v", portString)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	shutdownServerGracefully(server, logger)
}

func shutdownServerGracefully(server *http.Server, logger *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	logger.Println("Received", sig, "signal... Gracefully shutdown")

	timeOutContext, cancelTimeoutContext := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelTimeoutContext()

	go func() {
		<-timeOutContext.Done()
		if timeOutContext.Err() == context.DeadlineExceeded {
			logger.Println("graceful shutdown timed out... forcing exit.")
		}
	}()

	err := server.Shutdown(timeOutContext)
	if err != nil {
		logger.Fatal(err)
	}
}
