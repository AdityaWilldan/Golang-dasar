package main

import "fmt"

func main() {

	nama := "wildan"

	if nama == "wildan" {
		fmt.Println("hallo", nama)
	} else if nama == "wildan aditia" {
		fmt.Println("hallo", nama)
	} else {
		fmt.Println(nil)
	}

	if lenght := len(nama); lenght > 10 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
