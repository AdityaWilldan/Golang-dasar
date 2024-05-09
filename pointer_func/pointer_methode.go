package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "MR." + man.Name
}

func main() {

	wildan := Man{"jnck"}
	wildan.married()
	fmt.Println(wildan.Name)

}
