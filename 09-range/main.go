package main

import "fmt"

func main() {
	fmt.Println("Range")

	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids

	for i, id := range ids {
		fmt.Println("ID (", i, "):", id)
	}

	// Use an underscore if you are not using the index
	// for _, id := range ids {
	// 	fmt.Println("ID (", i, "):", id)
	// }

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Range with map

	coments := map[string]string{"Bob": "Hi I'm Bob", "Rose": "Hi I'm Rose"}

	for k, v := range coments {
		fmt.Println("( Key:", k, "): Value: '", v, "'")
	}

}
