// Metody v programovacím jazyku Go
//
// - metoda, která se snaží modifikovat příjemce
// - přístup k prvkům příjemce s využitím dereference

package main

import (
	"fmt"
)

// deklarace uživatelského typu
type Rectangle struct {
	width, height float64
}

// implementace metody
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// metoda, která se snaží modifikovat příjemce
func (rectangle *Rectangle) enlarge(horizontal float64, vertical float64) {
	// přístup k prvkům příjemce s využitím dereference
	(*rectangle).width += horizontal
	(*rectangle).height += vertical
}

func main() {
	// konstrukce obdélníka
	rect := Rectangle{width: 2, height: 3}

	// původní obdélník
	fmt.Println(rect)

	// pokus o změnu velikosti
	rect.enlarge(1, 2)

	// nový(?) obdélník
	fmt.Println(rect)
}
