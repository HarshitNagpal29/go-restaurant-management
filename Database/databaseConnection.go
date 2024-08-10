package Database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Println("Connecting to MongoDB", MongoDb)

	Client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return Client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(Client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("restaurant").Collection(collectionName)

	return collection
}
