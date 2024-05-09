package main

import "fmt"

//variadic function
func number(numbers ...int) int {
	total := 0
	for _, nilai := range numbers {
		total += nilai
	}
	return total

}

func main() {
	//variadic function
	hasil := number(10, 90, 30, 40, 50)
	fmt.Println(hasil)

	// slice paramete variadic
	slice := []int{10, 10, 10, 10}
	hasil = number(slice...)
	fmt.Println(hasil)
}
