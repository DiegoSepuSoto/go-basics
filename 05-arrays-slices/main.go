package main

import "fmt"

func main() {
	fmt.Println("Arrays and slices")
	// Arrays
	// Must have FIXED size
	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr)

	// Slices
	// Variable size
	colorSlice := []string{"Blue", "Yellow", "White"}
	fmt.Println(colorSlice)
	fmt.Println("Length", len(colorSlice))
	fmt.Println("[1:2]", colorSlice[1:2])
}
