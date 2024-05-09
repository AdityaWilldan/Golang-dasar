package main

import (
	"fmt"
	"reflect"
)

type Sampel struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == " " {
				return false
			}
		}
	}
	return true
}

func main() {

	sampel := Sampel{"Wildan"}
	sampelType := reflect.TypeOf(sampel)
	strukturField := sampelType.Field(0)
	required := strukturField.Tag

	fmt.Println(required.Get("required"))
	fmt.Println(required.Get("max"))

	sampel.Name = " "
	fmt.Println(IsValid(sampel))

}
