package main

import (
	"context"
	"log"
	"time"

	"github.com/amir-amirov/go-url-shortener/internal/db"
	"github.com/amir-amirov/go-url-shortener/internal/server"
	"github.com/amir-amirov/go-url-shortener/internal/shorten"
	"github.com/amir-amirov/go-url-shortener/internal/storage/shortening"
)

func main() {

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer dbCancel()

	const mongoDSN = "mongodb+srv://mongodb:mongodb@mycluster.fuoiz.mongodb.net/"
	const databaseName = "urlshortener"

	mgoClient, err := db.Connect(dbCtx, mongoDSN)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := mgoClient.Client().Disconnect(dbCtx); err != nil {
			log.Printf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	shorteningStorage := shortening.NewMongoDB(mgoClient.Client())
	shortener := shorten.NewService(shorteningStorage)
	server := server.New(shortener)

	server.Run(":8080")
}
