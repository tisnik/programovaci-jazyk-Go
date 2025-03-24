// Řetězce v programovacím jazyku Go
//
// - prázdný řetězec
// - řetězce s řídicími znaky
// - řetězce s Unicode znaky

package main

import "fmt"

func main() {
	// prázdný řetězec
	var s1 string

	// prázdný řetězec
	var s2 string = ""

	// řetězec s řídicími znaky
	var s3 string = "Hello\nworld!"

	// řetězec s Unicode znaky
	var s4 string = "ěščřžýáíéů"

	// řetězec s Unicode znaky
	var s5 string = "шщэюя"

	// tisk všech výše definovaných řetězců
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
