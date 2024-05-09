package halah

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

// func MyFunc(name string) string {
// 	return "hello " + name
// }
