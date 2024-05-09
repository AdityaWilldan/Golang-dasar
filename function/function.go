package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

//function parameter
func hitung(a int, b int) {
	fmt.Println(a * b)
}

//function parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

//reuturn function
func getHello(name string) string {
	if name == "" {
		return "hello broh"
	} else {
		return "Hello " + name
	}
}

//returning multiple value
func fullName() (string, string, string) {
	return "jancuk", "pristel", "cokk"
}

func main() {
	//function parameter
	sayHello()

	for i := 0; i < 5; i++ {
		sayHello()
	}
	//function parameter
	hitung(12, 32)

	//function parameter
	sayHelloTo("Wildan", "Aditia")

	//reuturn function
	resutl := getHello("jancuk")
	fmt.Println(resutl)
	fmt.Println(getHello(""))

	//returning multiple value
	namaDepan, _, namaBelakang := fullName()
	fmt.Println(namaDepan, namaBelakang)
}
