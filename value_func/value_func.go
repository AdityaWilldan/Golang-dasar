package main

import "fmt"

func getGooBye(name string) string {
	return "sayonara " + name
}

func main() {

	bye := getGooBye
	fmt.Println(bye("Wildan"))

}
