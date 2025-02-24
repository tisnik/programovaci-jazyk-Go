// Řezy v programovacím jazyku Go
package main

import "fmt"

func main() {
	var s []int

	fmt.Printf("Type:  %T\n", s)

	fmt.Printf("Before modification: %v\n", s)

	s[0] = 100

	fmt.Printf("After modification:  %v\n", s)
}
