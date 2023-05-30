package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dshwetank1:<gakuccespif0>@cluster0.gwklm34.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// Taking the reference of MongoDB collection (IMP)
var collection *mongo.Collection

// connect with the MongoDB
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	// can use the context.Background() ass well here
	// context is required every time, whenever the connection is need to be established to some reference machine
	client,err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mogno DB connection succesful")


	collection = client.Database(dbName).Collection(collectionName)

	// reference of the collection
	fmt.Println("Collection reference is ready")
}
