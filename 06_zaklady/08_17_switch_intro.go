// Základy programovacího jazyka Go
//
// - prázdná konstrukce switch
// - umístění větve default v konstrukci switch

package main

import "fmt"

func main() {
	switch {
	}

	switch {
	default:
		fmt.Println("proč jsem vlastně použil switch?")
	}

	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default")
	}

	switch {
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	case true:
		fmt.Println("true")
	}
}
