package main

import "fmt"

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func (dog Dog) toString() string {
	return fmt.Sprintf("Dog name: %s, breed: %s, age: %d", dog.Name, dog.Breed, dog.Age)
}

func main() {
	dog := Dog{
		Name:  "Rex",
		Breed: "Lab",
		Age:   8,
	}

	fmt.Println(dog.toString())
}
