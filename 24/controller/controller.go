package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/thakkarnetram/modules/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	url string
	dbName="netflix"
	colName="watchlist"
)

// collection imp 
var collection *mongo.Collection

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	// Load environment variables after loading the .env file
	url = os.Getenv("ATLAS_URI")
	fmt.Println("URI ", url)

	// client option
	clientOption:=options.Client().ApplyURI(url)
	// connection
	client, err :=mongo.Connect(context.TODO(),clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection success ")
	// collection 
	collection = client.Database(dbName).Collection(colName)
	// collection instance 
	fmt.Println("Collection ref is ready  ",collection)
}

// mongo helpers
// insert 1 
func insertOneMovie(movie model.Netflix){
	inserted , err := collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie " , inserted.InsertedID)
}

// update 1 movie
func updateOneMovie (movieId string ) {
	id,err:=primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter:= bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	// updating
	res,err:=collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count " , res.ModifiedCount)
}

// delete 1 movie
func deleteOneMovie (movieId string){
	id,err:= primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id":id} 
	res,err := collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count " , res.DeletedCount)
}

// delete all record
func deleteAllMovie () int64 {
	// {{}} means everything will be selected
	res,err := collection.DeleteMany(context.Background(),bson.D{{}},nil) // nil is optional
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count of movies deleted  ", res.DeletedCount)
	return res.DeletedCount
}

// get all movies
func getAllMovies() []primitive.M {
	cursor , err := collection.Find(context.Background(),bson.D{{}})
	// cursor is a type of mongodb object
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()){ 
		var movie bson.M 
		err:=cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// router controllers
// getting movies 
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	allMovies:= getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// creating movies 
func CreateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ =json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// isMovie Watched
func MarkedWatch(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application-json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	// id 
	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

// delete one  movie 
func DeleteMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application-json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	// id 
	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// delete all movies
func DeleteMovies(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application-json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	// id 
	count:= deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}