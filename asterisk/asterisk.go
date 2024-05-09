package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var Address1 Address = Address{"Subang", "Jawa barat", "Indonesia"}
	var Address2 *Address = &Address1

	Address2.City = "surabaya"
	fmt.Println(Address1)
	fmt.Println(Address2)

	*Address2 = Address{"Jakarta", "DKI", "Indonesia"}
	fmt.Println(Address2)
	fmt.Println(Address1)
}
