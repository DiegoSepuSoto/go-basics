package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Key-value pair
	emails := make(map[string]string)

	// Assing kv
	emails["Bob"] = "bob@gmail.com"
	emails["Rose"] = "rose@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Mike")
	fmt.Println(emails)

	// Declare & assign
	coments := map[string]string{"Bob": "Hi I'm Bob", "Rose": "Hi I'm Rose"}
	fmt.Println(coments)
}
