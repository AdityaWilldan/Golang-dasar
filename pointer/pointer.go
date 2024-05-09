package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

// pointer func

func changeCountryToIndonesia(address *Address) {
	address.Country = "indonesia"
}

func main() {
	// Address1 := Address{"Subang", "Jawa barat", "Indonesia"}
	// Address2 := Address1

	// Address2.City = "bandung"
	// fmt.Println(Address1)
	// fmt.Println(Address2)

	var Address1 Address = Address{"Subang", "Jawa barat", "Indonesia"}
	var Address2 *Address = &Address1
	var Address3 = &Address1

	Address2.City = "surabaya"

	*Address2 = Address{"Jakarta", "DKI Jakarta", "indonesia"}
	fmt.Println(Address1)
	fmt.Println(Address2)
	fmt.Println(Address3)

	// membuat ponter kosong menggunakan new():
	Address4 := new(Address)
	Address4.City = "Sukabumi"
	fmt.Println(Address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jabar",
		Country:  " ",
	}
	var alamatPointer *Address = &alamat
	changeCountryToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}
