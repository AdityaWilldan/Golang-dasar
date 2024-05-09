package main

import "fmt"

type Hashname interface {
	Getname() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHello(hashname Hashname) {
	fmt.Println("Hello", hashname.Getname())
}

func (person Person) Getname() string {
	return person.Name
}

func (animal Animal) Getname() string {
	return animal.Name
}

func main() {
	var wildan Person
	wildan.Name = "Wildan"

	sayHello(wildan)

	var kucing Animal
	kucing.Name = "kucing"

	sayHello(kucing)

}
