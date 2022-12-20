package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var person = new(Person)

	person.name = "hukangkang"
	person.age = 12

	fmt.Printf("%T", *person)
}
