// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch s výrazy ve větvích case.

package main

import "fmt"

func classify(x int) string {
	switch {
	case x == 0:
		return "nula"
	case x%2 == 0:
		return "sudé číslo"
	case x%2 == 1:
		return "liché číslo"
	default:
		return "k tomu nemůže dojít"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x))
	}
}
