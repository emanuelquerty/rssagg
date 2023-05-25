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
	"github.com/emanuelquerty/rssagg/internal/routers"
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
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := routers.NewRouter()
	v1Router.DBConn = dbConn
	v1Router.Logger = logger
	router.Mount("/v1", v1Router.Route())

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
