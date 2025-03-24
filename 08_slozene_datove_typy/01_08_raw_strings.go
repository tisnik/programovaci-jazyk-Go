// Řetězce v programovacím jazyku Go
//
// - deklarace a použití takzvaných surových řetězců
//   (raw strings)

package main

import "fmt"

func main() {
	// běžný řetězec s řídicími znaky
	var s1 string = "Hello\nworld!\n"

	// surový řetězec (bez řídicích znaků)
	var s2 string = `Hello\nworld!\n`

	fmt.Println(s1)
	fmt.Println(s2)
}
