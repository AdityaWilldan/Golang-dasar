package main

import (
	"fmt"
	"strconv"
)

func main() {
	// mengubah string menjadi boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}
	// mengubah string menjadi integer
	number, err := strconv.ParseInt("1200000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// mengubah integer ke string
	nilai := strconv.FormatInt(1000000, 10)
	fmt.Println(nilai)

	nilaiINT, _ := strconv.Atoi("2000000")
	fmt.Println(nilaiINT)
}
