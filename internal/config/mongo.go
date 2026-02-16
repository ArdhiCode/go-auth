package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() *mongo.Client {
	uri := GetEnv("MONGO_URI", "mongodb://localhost:27017")
	db := GetEnv("MONGO_DB", "web-arrahmat-cbt")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB Connected to ", db)
	return client
}
