package main

import "fmt"

// email := "diego@sepu.cl"
// syntax error: non-declaration statement outside function body

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using 'var'
	var name string = "Diego Sepúlveda"
	// var name = "Diego Sepúlveda"
	// const name = "Diego Sepúlveda"
	fmt.Println("Name: " + name)
	var age int = 22
	fmt.Println("Age:", age)
	// fmt.Printf("Age: %d\n", age)
	fmt.Printf("Name-Type: %T - Age-Type: %T\n", name, age)
	email := "diego@sepu.cl"
	fmt.Println("Email: " + email)
	favLang1, favLang2 := "JS", "BASIC"
	fmt.Println("Favorite Programming Language 1: " + favLang1)
	fmt.Println("Favorite Programming Language 2: " + favLang2)
}
