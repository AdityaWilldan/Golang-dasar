package main

import "fmt"

func Factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * Factorial(value-1)
	}
}

func main() {

	fmt.Println(Factorial(10))

}
