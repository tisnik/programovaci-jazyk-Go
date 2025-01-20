// Základy programovacího jazyka Go
//
// - řídicí konstrukce if-else

package main

import "fmt"

func classify_char(c rune) string {
	if c >= 'a' && c <= 'z' {
		return "male pismeno"
	} else if c >= 'A' && c <= 'Z' {
		return "velke pismeno"
	} else {
		return "neco jineho"
	}
}

func main() {
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if !true {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	fmt.Println(classify_char('a'))
	fmt.Println(classify_char('Z'))
	fmt.Println(classify_char('?'))
}
