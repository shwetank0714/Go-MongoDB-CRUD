package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/shwetank0714/mongodbapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dshwetank1:gakuccespif0@cluster0.gwklm34.mongodb.net/?retryWrites=true&w=majority"
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
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mogno DB connection succesful")

	collection = client.Database(dbName).Collection(collectionName)

	// reference of the collection
	fmt.Println("Collection reference is ready")
}

// Mongo DB Helpers - file

// Insert a record ( Add a data in Mongo)
func insertMovie(movie *model.Netflix) error{
	insertedResult, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
		return err
	}

	movie.ID = primitive.NewObjectID()

	log.Println("INSERTED RESULT DATA", insertedResult)
	log.Println("Movie inserted succesfully in DB, ID: ", insertedResult.InsertedID)

	return nil
}

// Update the record in DB
func updateMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_watched": true}}
	// gives the result - how many records are updated
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)

}

// Delete single record from MongoDB
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteResult, _ := collection.DeleteOne(context.Background(), filter)

	fmt.Println("Data deleted: modified deleted count", deleteResult.DeletedCount)
}

// Delete All records from MongoDB
func deleteAllMovies() int64 {
	// {} inside bson.M will take all the values that are present in DB
	filter := bson.D{{}}
	deleteResult, _ := collection.DeleteMany(context.Background(), filter, nil)

	fmt.Println("deleted All data count is: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

// Get All data from MongoDB
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

// CONTROLLERS - File ----------------------------------------------------------------------

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil{
		log.Fatal(err)
	}

	if err := insertMovie(&movie); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie)
}

func MakrMovieAsWatched(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
