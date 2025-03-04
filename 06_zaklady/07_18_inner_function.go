// Základy programovacího jazyka Go
//
// - deklarace funkce uvnitř jiné funkce
// - zavolání této funkce s předáním argumentů

package main

import (
	"fmt"
	"math"
)

// Výpočet přepony pravoúhlého trojúhelníku
func hypot(a, b float64) float64 {
	// vnitřní funkce
	square := func(x float64) float64 {
		return x * x
	}

	// dvojí zavolání vnitřní funkce
	c := math.Sqrt(square(a) + square(b))
	return c
}

func main() {
	// zavolání funkce s předáním parametrů
	fmt.Println(hypot(3.0, 4.0))
}
