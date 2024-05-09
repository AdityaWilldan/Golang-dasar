package main

import "fmt"

func main() {
	type NoKtp string
	type mahasiswa bool

	var NoKtpWildan NoKtp = "27838197418491"
	var status mahasiswa = true

	fmt.Println(NoKtpWildan)
	fmt.Println(status)
}
