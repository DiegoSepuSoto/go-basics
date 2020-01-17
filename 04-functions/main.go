package main

import "fmt"

func greeting(name string) string {
	return "Greetings, " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Functions")
	fmt.Println(greeting("Diego"))
	fmt.Println("1 + 1 =", getSum(1, 1))
}
