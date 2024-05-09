package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("w([a-z])l")

	fmt.Println(regex.MatchString("wildan"))
	fmt.Println(regex.MatchString("danwil"))
	fmt.Println(regex.MatchString("nadliw"))

	search := regex.FindAllString("wildan danwil nadliw nidlaw liwdan dildan danliw", 10)

	fmt.Println(search)
}
