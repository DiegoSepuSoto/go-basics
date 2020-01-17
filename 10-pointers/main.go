package main

import "fmt"

func main() {

	fmt.Println("Pointers")

	a := 5
	b := &a

	fmt.Println("b:", b)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", *b)
	a += 10
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", *b)
	*b = 20
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", *b)
}
