package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Wildan aditia", "Wildan"))
	fmt.Println(strings.Split("Wildan aditia", " "))
	fmt.Println(strings.ToLower("WILDAN ADITIA"))
	fmt.Println(strings.ToUpper("wildan aditia"))
	fmt.Println(strings.Trim("   wildan aditia    ", " "))
	fmt.Println(strings.ReplaceAll("kiwil kiwil kiwil kiwil", "kiwil", "wildan"))
}
