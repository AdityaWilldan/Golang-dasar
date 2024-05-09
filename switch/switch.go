package main

import "fmt"

func main() {

	name := "wildan"

	switch name {
	case "wildan":
		fmt.Println("hallo", name)
	case "wildan aditia":
		fmt.Println("hai!", name)
	default:
		fmt.Println("hei,boleh kenalan?")
	}

	//switch case short statement

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama sudah benar!")
	// case false:
	// 	fmt.Println("nama terlalu panjang")
	// }

	length := len(name)
	switch {
	case length > 15:
		fmt.Println("nama terlalu panjang")
	case length > 10:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}
