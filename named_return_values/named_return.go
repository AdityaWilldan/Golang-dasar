package main

import "fmt"

func getCompleteName() (firstname, lastname, status string) {
	firstname = "Wildan"
	lastname = "Aditia"
	status = "Mahasiswa"

	return
}

func main() {

	fmt.Println(getCompleteName())
}
