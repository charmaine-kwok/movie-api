package initializers

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *mongo.Database
var CRED_DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	cred_db := os.Getenv("CRED_URL")

	CRED_DB, err = gorm.Open(postgres.Open(cred_db), &gorm.Config{})
	if err != nil {
		panic("failed to connect credential database")
	}
	_ = CRED_DB

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}

	// Get the database and collection
	DB = client.Database("Movies")
}
