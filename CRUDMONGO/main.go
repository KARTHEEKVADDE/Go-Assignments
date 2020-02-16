//CRUD GO-MONGO
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"reflect"
	"time"
)

var client *mongo.Client

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":12345", router)
}

// CreateUserHandler creates a new user
func CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("Kartheek").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// Call the InsertOne() method by passing BSON
	result, _ := collection.InsertOne(ctx, user)
	fmt.Println("CreateUser Result TYPE:", reflect.TypeOf(result))
	json.NewEncoder(response).Encode(result)
}

// GetUserHandler gets an existing user
func GetUserHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("Kartheek").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Call the FindOne() method by passing BSON
	result := collection.FindOne(ctx, User{ID: id})
	_ = result.Decode(&user)
	fmt.Println("GetUser Result TYPE:", reflect.TypeOf(result))
	json.NewEncoder(response).Encode(user)
}

// GetUsersHandler gets all existing users
func GetUsersHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []User
	collection := client.Database("Kartheek").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Call the Find() method by passing BSON
	result, _ := collection.Find(ctx, bson.M{})
	defer result.Close(ctx)

	for result.Next(ctx) {
		var user User
		result.Decode(&user)
		users = append(users, user)
	}
	fmt.Println("GetUsers Result TYPE:", reflect.TypeOf(result))
	json.NewEncoder(response).Encode(users)
}

// DeleteUserHandler deletes an existing user
func DeleteUserHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("Kartheek").Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Call the DeleteOne() method by passing BSON
	result, _ := collection.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println("DeleteUser Result TYPE:", reflect.TypeOf(result))
	json.NewEncoder(response).Encode(result)
}

// UpdateUserHandler updates an existing user
func UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("Kartheek").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Call the UpdateOne() method by passing BSON
	result, _ := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"firstname", user.Firstname}, {"lastname", user.Lastname}}},
		},
	)
	fmt.Println("UpdateUser Result TYPE:", reflect.TypeOf(result))
	json.NewEncoder(response).Encode(result)
}
