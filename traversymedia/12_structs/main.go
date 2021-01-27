package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	fmt.Println(p.firstName)
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
	fmt.Println(p.firstName)
}

func main() {
	person1 := Person{firstName: "Pierce", lastName: "Faraone", city: "Singapore", gender: "m", age: 26}
	person2 := Person{firstName: "Angad", lastName: "Sethi", city: "London", gender: "m", age: 26}
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person1.hasBirthday()
	person2.hasBirthday()
	person2.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	// fmt.Println(person1.greet())
	// fmt.Println(person1.greet())
	// person1.hasBirthday()
	// fmt.Println(person1.greet())
}
