package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {

	blockUser := func(name string) bool {
		return name == "admin"

	}

	registerUser("admin", blockUser)
	registerUser("wildan", blockUser)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("adita", func(name string) bool {
		return name == "root"
	})

}
