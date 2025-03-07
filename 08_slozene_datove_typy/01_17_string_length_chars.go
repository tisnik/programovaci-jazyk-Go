// Řetězce v programovacím jazyku Go
//
// - zjištění délky řetězce v runách

package main

import (
	"fmt"
)

func main() {
	var s1 string
	var s2 string = ""
	var s3 string = "Hello world!"
	var s4 string = "ěščřžýáíéů"
	var s5 string = "шщэюя"

	fmt.Println(len([]rune(s1)))
	fmt.Println(len([]rune(s2)))
	fmt.Println(len([]rune(s3)))
	fmt.Println(len([]rune(s4)))
	fmt.Println(len([]rune(s5)))
}
