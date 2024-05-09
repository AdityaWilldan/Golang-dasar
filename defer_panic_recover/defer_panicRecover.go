package main

import (
	"fmt"
)

// defer
func logging() {
	fmt.Println("selesai memanggil function")
}

// defer
func RunApp(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 15 / value
	fmt.Println(result)
}

// panic & recover
func RunServer() {
	fmt.Println("running server on port 8080")
	message := recover()
	if message != nil {
		fmt.Println("Server terkendala", message)
	}
}

// panic & recover
func check(error bool) {
	defer RunServer()
	if error {
		panic("ERROR!")
	}
	fmt.Println("server berjalan")

}

func main() {

	RunApp(5)

	fmt.Print("\n")

	check(false)
	fmt.Println("wildan")

}
