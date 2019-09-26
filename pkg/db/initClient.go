package db

import (
	"context"
	"github.com/Blockchainpartner/hagentrust/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var db *mongo.Database

// InitClient initializes a DB client object
func InitClient() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(configs.MongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	// create DB client object
	db = client.Database(configs.MongodbDBName)
}