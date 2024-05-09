package main

import (
	"fmt"
	"time"
)

func main() {
	waktu := time.Now()
	fmt.Println(waktu.Local())
	fmt.Println(waktu.Month())
	fmt.Println(waktu.Hour())
	fmt.Println(waktu.Date())
	fmt.Println(waktu.Day())
	fmt.Println(waktu.Year())
	fmt.Println(waktu.Minute())
	fmt.Println(waktu.Second())
	fmt.Println(waktu.Nanosecond())

	utc := time.Date(2024, 8, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2021-10-01")
	fmt.Println(parse)
}
