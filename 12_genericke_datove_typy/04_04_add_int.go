// Generické datové typy v Go
//
// - negenerická (běžná) funkce pro součet
//   dvou hodnot typu celé číslo

package main

import "fmt"

// negenerická (běžná) funkce pro součet
// dvou hodnot typu celé číslo
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
