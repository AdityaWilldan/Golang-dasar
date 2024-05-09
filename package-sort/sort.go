package main

import (
	"fmt"
	"sort"
)

type user struct {
	name string
	age  int
}

type UserSlice []user

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []user{
		{"Wildan", 200},
		{"Kiwil", 35},
		{"Aditia", 41},
	}
	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
