package main

import (
	"fmt"
)

func random() interface{} {
	return 21
}

func main() {

	result := random()

	// resultString := result.(string)
	// fmt.Println(resultString)

	//gunakan switch untuk bisa menetukan hasil tipe data
	switch hasil := result.(type) {
	case string:
		fmt.Println("ini string", hasil)
	case int:
		fmt.Println("integer coy", hasil)
	case bool:
		fmt.Println("boolean", hasil)
	default:
		fmt.Println("Unknown")

	}

}
