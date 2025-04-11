// Generické datové typy v Go
//
// - negenerická (bežná) funkce pro součet
//   dvou hodnot typu celé číslo

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
