// Základy programovacího jazyka Go
//
// - zkrácené vyhodnocení operátorů && a ||

package main

import "fmt"

func f1() bool {
	println("f1")
	return true
}

func f2() bool {
	println("f2")
	return false
}

func f3() bool {
	println("f2")
	return false
}

func main() {
	// zkrácené vyhodnocení operátorů && a || má vliv na to,
	// které funkce f1() a f2() se budou volat
	fmt.Printf("short circuit &&: %v\n", f1() && f2())
	fmt.Printf("short circuit ||: %v\n", f1() || f2())
	fmt.Printf("short circuit &&: %v\n", f2() && f3())
	fmt.Printf("short circuit ||: %v\n", f2() || f3())
}
