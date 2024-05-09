package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("wildan")
	data.PushBack("kiwil")
	data.PushBack("aditia")
	data.PushFront("jnck")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for nilai := data.Front(); nilai != nil; nilai = nilai.Next() {
		fmt.Println(nilai.Value)
	}
	for nilai := data.Back(); nilai != nil; nilai = nilai.Prev() {
		fmt.Println(nilai.Value)
	}

}
