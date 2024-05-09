package main

import "fmt"

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {

	test := ups(2)
	fmt.Println(test)

	var coba interface{} = ups(3)
	fmt.Println(coba)

}
