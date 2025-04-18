// Generické datové typy v Go
//
// - deklarace interních proměnných s využitím
//   generického datového typu

package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric](x T, y T) T {
	// deklarace interních proměnných s využitím
	// generického datového typu
	var first T = x
	var second T = y
	return first + second
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3.14, 6.28))
}
