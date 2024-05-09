package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian tidak bisa di bagi dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, error := Pembagian(100, 10)
	if error == nil {
		fmt.Println("hasilnya adalah: ", hasil)
	} else {
		fmt.Println("error boss", error.Error())
	}

}
