//Time Incrementor
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Request Body Model - Response Body Model
type RequestBody struct {
	InitialTime string `json:"initialTime"`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting the application...")
	router.HandleFunc("/time", TimeHandler).Methods("POST")
	http.ListenAndServe(":12345", router)
}

func TimeHandler(response http.ResponseWriter, request *http.Request) {
	// Initialize Default Time
	timeString := RequestBody{InitialTime: "2020-02-16T12:40:02"}

	// Parse JSON Request Body (Take from Request)
	_ = json.NewDecoder(request.Body).Decode(&timeString)
	// Parse the Time based on a layout
	parsedTime, _ := time.Parse("2006-01-02T15:04:05", timeString.InitialTime)
	fmt.Println(parsedTime)
	// Increment the Time (Incrementing for 12 Hours)
	AdvancedTime := parsedTime.Add(time.Hour * 12)
	fmt.Println(AdvancedTime)
	timeString.InitialTime = AdvancedTime.Format("2006-01-02T15:04:05")

	// Marshal or Convert time object back to json
	timeResponse, _ := json.Marshal(timeString)

	// Set Content-Type Header
	response.Header().Set("Content-Type", "application/json")
	// Write To JSON Response Body
	response.Write(timeResponse)
}
