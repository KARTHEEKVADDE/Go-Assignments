//PERLIN NOISE 2
package main

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	//"math/rand"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 100
	X           = 50 // X-Co-ordinate
	Y           = 50 // Y-Co-ordinate
	Z           = 50 // Z-Co-ordinate
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting the Application...")
	router.HandleFunc("/1d-noise", PerlinNoise1DHandler).Methods("GET") //Generates 1D Noise
	router.HandleFunc("/2d-noise", PerlinNoise2DHandler).Methods("GET") //Generates 2D Noise
	router.HandleFunc("/3d-noise", PerlinNoise3DHandler).Methods("GET") //Generates 3D Noise
	http.ListenAndServe(":12345", router)
}

func PerlinNoise1DHandler(response http.ResponseWriter, request *http.Request) {
	result := []float64{}
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < X; x++ {
		fmt.Printf("%0.0f\t%0.4f\n", x, p.Noise1D(x/10))
		result = append(result, p.Noise1D(x/10))
	}
	// Set Content-Type Header
	response.Header().Set("Content-Type", "application/json")
	// Write to JSON Response Body
	s := fmt.Sprintf("%s %f %d", "Generating 1D Noise:", result, len(result))
	response.Write([]byte(s))
}

func PerlinNoise2DHandler(response http.ResponseWriter, request *http.Request) {
	result := []float64{}
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < X; x++ {
		for y := 0.; y < Y; y++ {
			fmt.Printf("%0.0f\t%0.0f\t%0.4f\n", x, y, p.Noise2D(x/10, y/10))
			result = append(result, p.Noise2D(x/10, y/10))
		}
	}
	// Set Content-Type Header
	response.Header().Set("Content-Type", "application/json")
	// Write to JSON Response Body
	s := fmt.Sprintf("%s %f %d", "Generating 2D Noise:", result, len(result))
	response.Write([]byte(s))
}

func PerlinNoise3DHandler(response http.ResponseWriter, request *http.Request) {
	result := []float64{}
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < X; x++ {
		for y := 0.; y < Y; y++ {
			for z := 0.; z < Z; z++ {
				fmt.Printf("%0.0f\t%0.0f\t%0.0f\t%0.4f\n", x, y, z, p.Noise3D(x/10, y/10, z/10))
				result = append(result, p.Noise3D(x/10, y/10, z/10))
			}
		}
	}
	// Set Content-Type Header
	response.Header().Set("Content-Type", "application/json")
	// Write to JSON Response Body
	s := fmt.Sprintf("%s %f %d", "Generating 3D Noise:", result, len(result))
	response.Write([]byte(s))
}
