package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"
)

func Config(dbURL string) *sql.DB {
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to use data source name:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appSignal := make(chan os.Signal, 1)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		sig := <-appSignal
		log.Printf("Received %v signal... database resources released", sig)
		cancel()
	}()

	ctx, stop := context.WithTimeout(ctx, 1*time.Second)
	defer stop()

	if err := dbConn.PingContext(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Connection to database has been established")

	return dbConn
}
