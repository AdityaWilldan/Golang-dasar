package main

import (
	"fmt"
	"latihan_golang/coba"
	"latihan_golang/halah"
)

func bubblesort(arr []int) {
	var temp int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {

	arr := []int{1, 3, 2, 5, 4, 10, 13, 12, 8, 6, 7, 9}
	bubblesort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("hello world")

	name := "John"
	message := coba.MyFunc(name)
	fmt.Println(message)

	blockUser := func(name string) bool {
		return name == "admin"

	}

	halah.RegisterUser("admin", blockUser)
	halah.RegisterUser("wildan", blockUser)

	// nama := "anyink"
	// pesan := halah.MyFunc(nama)
	// fmt.Println(pesan)
}
