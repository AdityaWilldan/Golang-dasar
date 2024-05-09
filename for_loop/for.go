package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	fmt.Print("\n")

	for nilai := 1; nilai <= 10; nilai++ {
		fmt.Println("nilai ke-", nilai)
	}

	fmt.Print("\n")
	// for range
	names := []string{"wildan", "aditia", "uhuuyy"}

	for i, name := range names {
		fmt.Println("index", i, "=", name)
	}

	person := make(map[string]string)
	person["title"] = "Progammer"
	person["status"] = "Mahasiswa"
	person["gender"] = "male"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
