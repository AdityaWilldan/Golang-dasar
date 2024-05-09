package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args
	fmt.Println("arugument : ")
	fmt.Println(argument)

	name, error := os.Hostname()
	if error == nil {
		fmt.Println("hostname : ", name)
	} else {
		fmt.Println("error", error.Error())
	}

	username := os.Getenv("APP_USERNAME")
	passsword := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(passsword)

}
