package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	NameFilteRed := filter(name)
	fmt.Println("Hallo", NameFilteRed)
}

func spamFiltered(name string) string {
	if name == "Anjing" || name == "Asu" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Wildan aditia", spamFiltered)

	filter := spamFiltered
	sayHelloWithFilter("Anjing", filter)

}
