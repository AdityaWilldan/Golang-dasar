package main

import "fmt"

func main() {

	nama := [...]string{
		"wildan",
		"aditia",
		"bambang",
		"budi",
		"nugraha",
		"asep spakbor"}
	slice := nama[4:6]

	fmt.Println(slice)

	bulan := [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"december",
	}

	slice1 := bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// // bulan[5] = "diubah"
	slice2 := bulan[11:]
	fmt.Println(slice2)

	//append
	slice3 := append(slice2, "libur cooy")
	slice3[0] = "upss"
	fmt.Println(slice3)
	fmt.Println(bulan)

	//make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Wildan"
	newSlice[1] = "Aditia"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	// slice4 := newSlice[5:]
	// slice4[1] = "jancuk"
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(newSlice))

}
