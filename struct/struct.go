package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Status        bool
}

//struct function / methode
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	var Wildan Customer

	Wildan.Name = "wildan aditia"
	Wildan.Address = "Subang"
	Wildan.Age = 21
	fmt.Println(Wildan)

	Joko := Customer{
		Name:    "Djoko suroko",
		Address: "Jawa Timur",
		Age:     31,
	}
	fmt.Println(Joko)

	Arif := Customer{"Arif Suroto", "Magelang", 41}
	fmt.Println(Arif)

	//struct function / methode
	Wildan.sayHello("Joko")

}
