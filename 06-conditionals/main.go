package main

import "fmt"

func main() {
	fmt.Println("Conditionals")
	x := 50
	y := 10

	// if

	if x < y {
		fmt.Println("X is less than Y")
	} else {
		fmt.Println("X isn't less than Y")
	}

	color := "red"

	// else if

	if color == "red" {
		fmt.Println("The color is red")
	} else if color == "pink" {
		fmt.Println("The color is pink")
	} else {
		fmt.Println("The color is neither red nor pink")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "pink":
		fmt.Println("The color is pink")
	default:
		fmt.Println("The color is neither red nor pink")
	}
}
