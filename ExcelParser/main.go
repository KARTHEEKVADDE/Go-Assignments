//Excel Parser
package main

import (
	"fmt"
)

var initialString string
var columns, rows int

var result = []string{}
var count = 0

func main() {
	fmt.Println("A SIMPLE GO-EXCEL PARSER")
Input:
	fmt.Println("Please Enter Starting Column String, Nō of Columns, Nō of Rows ")
	fmt.Scanf("%s %d %d", &initialString, &columns, &rows)
	totalCount := columns * rows

	if len(initialString) > 2 {
		fmt.Println("Please provide one(1) or two(2) character string! :)")
		goto Input
	}
	fmt.Println("Columns:", columns, "Rows:", rows, "Total Count:", totalCount)
	fmt.Println("-----------------------")

	if len(initialString) == 1 {
		start := int(initialString[0])
		second := 64
		solve(start, second, totalCount)
	} else if len(initialString) == 2 {
		start := int(initialString[1])
		second := int(initialString[0])
		solve(start, second, totalCount)
	}
	print() // Result in Tabular Format(Row*Column)
}

func solve(start, second, totalCount int) {
	for countX := 0; countX < totalCount; countX++ {
		if start == 91 {
			start = 65
			second++
		}
		if second < 65 {
			result = append(result, string(start))
			fmt.Print(string(start), " ")
		} else if second > 90 {
			fmt.Println("\n :) Limit Reached :)\n")
			break
		} else {
			result = append(result, string(second)+string(start))
			fmt.Print(string(second), string(start), " ")
		}
		start++
	}
}

func print() {
	fmt.Println("\nStart of the Result! :)")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if count < len(result) {
				fmt.Printf("|\t%s\t", result[count])
				count++
			} else {
				goto Exit
			}
		}
		fmt.Println("|")
	}
Exit:
	fmt.Println("End of the Result! :)")
}
