package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/decipher07/mongoapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/netflix"

/* Database Name */
const dbName = "netflix"

/* Collection Name */
const colName = "watchlist"

/* Collection Variable to be used by every other file for configuration - This helps in not making a connection everytime for the request */
var collection *mongo.Collection

/* Connect function to database */
func init() {

	/* Client Option */
	clientOption := options.Client().ApplyURI(connectionString)

	/* Connect to Mongodb */
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection Successful")

	/* Assigning the collection variable */
	collection = client.Database(dbName).Collection(colName)

	/* Collection Instance */
	fmt.Println("Collection Instance is ready!")
}

/* Insert a data */
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in Database with id : ", inserted.InsertedID)
}
