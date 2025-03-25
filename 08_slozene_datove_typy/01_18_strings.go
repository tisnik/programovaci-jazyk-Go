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
	s2 := ""

	// řetězec s řídicími znaky
	s3 := "Hello\nworld!"

	// řetězec s Unicode znaky
	s4 := "ěščřžýáíéů"

	// řetězec s Unicode znaky
	s5 := "шщэюя"

	// tisk všech výše definovaných řetězců
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
