//PERLIN NOISE 1
package main

import (
	"fmt"
	"github.com/aquilax/go-perlin"
)

const (
	alpha         = 2.
	beta          = 2.
	n             = 3
	seed  int64   = 100
	X     float64 = 100 // X-Co-ordinate
	Y     float64 = 100 // Y-Co-ordinate
)

// Result Channel for storing Noise Values
var result = make(chan float64)
var count = 1

func main() {
	fmt.Println("Starting the Application...")
	PerlinNoise2DGenerator()
	fmt.Scanln("")
}

func PerlinNoise2DGenerator() {
	//I have used two go routines here! You can have more! :)
	go halfResult(X/2, X, 0, Y)
	go halfResult(0, X/2, 0, Y)
	for i := range result {
		fmt.Println("Count=", count, "Noise=", i)
		count++
	}
}

func halfResult(startX, endX, startY, endY float64) {
	var p = perlin.NewPerlin(alpha, beta, n, seed)
	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			fmt.Printf("X=%0.0f\tY=%0.0f\tNoice=%0.4f\n", x, y, p.Noise2D(x/10, y/10))
			//result = append(result, p.Noise2D(x/10, y/10))
			result <- p.Noise2D(x/10, y/10)
		}
	}
}
