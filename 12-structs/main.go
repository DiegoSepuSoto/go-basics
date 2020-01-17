package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Methods
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years young"
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {

	fmt.Println("Structs")

	// person1 := Person{firstName: "Diego", lastName: "Sepúlveda", city: "Santiago", gender: "m", age: 22}

	person2 := Person{"Diego", "Sepúlveda", "Santiago", "m", 22}

	// fmt.Println(person1)
	// fmt.Println(person2)
	// fmt.Println(person2.firstName)
	// fmt.Println(person1.age)
	person2.hasBirthday()
	fmt.Println(person2.greet())

}
