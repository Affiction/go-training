package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func NewPerson(name string, age uint8) *Person {
	return &Person{name, age}
}

func main() {
	fmt.Println(Person{})
	fmt.Println(Person{Name: "Jake", Age: 18})
	fmt.Println(Person{"Jake", 18})
	fmt.Println(Person{Name: "Jake"})
	fmt.Println(NewPerson("Jake", 18))
}
