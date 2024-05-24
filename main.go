package main

import (
	"fmt"

	cust "github.com/marzelwidmer/gotdd/customer"
)

func main() {
	fmt.Println("Hello, World!")
	person := cust.Person{Name: "James", Age: 30}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
