package main

import (
	"fmt"

	"github.com/marzelwidmer/gotdd/customer"
	// cust "github.com/marzelwidmer/gotdd/customer"
	// "github.com/marzelwidmer/gotdd/customer"
)

func main() {
	fmt.Println("Hello, World!")
	person := customer.Person{Name: "James", Age: 30}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)

}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
