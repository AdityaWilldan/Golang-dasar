package main

import "fmt"

func main() {

	var nama [3]string

	nama[0] = "wildan"
	nama[1] = "aditia"
	nama[2] = "adit"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	nilai := [3]int{90, 95, 80}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	fmt.Println(len(nama))
	fmt.Println(len(nilai))

}
