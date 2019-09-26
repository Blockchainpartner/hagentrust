package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
)

var (
	// UserCollection points to User collection
	UserCollection *mongo.Collection
)

// InitCollections initializes collection objects and creates indexes
func InitCollections() {
	UserCollection = db.Collection("user")

	// Unique key on UserCollection
	userCollectionIndexView := UserCollection.Indexes()
	// Text version index on UserCollection
	textVersion := int32(3)
	indexName := "user_text_index"
	_, err := userCollectionIndexView.DropOne(context.Background(), indexName)
	if err != nil {
		log.Println(err)
	}
	opts := options.IndexOptions{
		Weights: bsonx.Doc{{
			Key:   "name",
			Value: bsonx.Int32(3),
		}},
		TextVersion: &textVersion,
		Name:        &indexName,
	}
	_, err = userCollectionIndexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bsonx.Doc{{
				Key:   "name",
				Value: bsonx.String("text"),
			}},
			Options: &opts,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
