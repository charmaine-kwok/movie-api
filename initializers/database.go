package initializers

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}

	// Get the database and collection
	DB = client.Database("Movies")
}
