package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "wildan",
		"address": "subang",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Print("\n")

	book := make(map[string]string)
	book["title"] = "golang e-book"
	book["Author"] = "Wildan"
	book["Penerbit"] = "Shonen Jump"
	fmt.Println(book)

	delete(book, "Penerbit")

	fmt.Println(book)

}
