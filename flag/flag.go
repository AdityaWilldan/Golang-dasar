package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "put your host")
	var user *string = flag.String("user", "root", "put your user")
	var password *string = flag.String("password", "root", "put your password")
	var number *int = flag.Int("number", 100, "put your number")

	flag.Parse()

	fmt.Println("host : ", *host)
	fmt.Println("user : ", *user)
	fmt.Println("password : ", *password)
	fmt.Println("number : ", *number)

}
