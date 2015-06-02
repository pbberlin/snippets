package main

import "fmt"

func main() {
	var response int
	fmt.Println("Y or y\n")
	// fmt.Scanf("%c", &response) //<--- here
	fmt.Scan(&response)
	fmt.Printf("reached next line\n")

	fmt.Println("\nxx", response)
}
