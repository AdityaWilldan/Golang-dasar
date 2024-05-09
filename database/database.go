package database

import "fmt"

var connection string

func init() {
	fmt.Println("function di panggil")
	connection = "mysql"
}

func GetDatabase() string {
	return connection
}
