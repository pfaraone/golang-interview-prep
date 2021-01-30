package main

import "fmt"

func main() {
	// can only pass by value or by pointer
	a := 5
	b := &ajj

	fmt.Println(a, b)
	fmt.Printf("%T \n", b)
}
