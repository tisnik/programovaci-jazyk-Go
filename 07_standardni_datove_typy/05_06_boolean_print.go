// Základní datové typy v jazyce Go
//
// - tisk hodnoty typu boolean
// - využití různých typů formátu

package main

import "fmt"

func main() {
	x := true
	y := false

	// tisk hodnoty typu boolean
	fmt.Println("%t format parameter:")
	fmt.Printf("x=%t\n", x)
	fmt.Printf("y=%t\n", y)

	fmt.Println()

	// tisk hodnoty typu boolean
	fmt.Println("%v format parameter:")
	fmt.Printf("x=%v\n", x)
	fmt.Printf("y=%v\n", y)

	fmt.Println()

	// tisk hodnoty typu boolean
	fmt.Println("left align:")
	fmt.Printf("| %-20t |\n", x)
	fmt.Printf("| %-20t |\n", y)

	fmt.Println()

	// tisk hodnoty typu boolean
	fmt.Println("right align:")
	fmt.Printf("| %20t |\n", x)
	fmt.Printf("| %20t |\n", y)
}
