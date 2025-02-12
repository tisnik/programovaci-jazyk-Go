// Základy programovacího jazyka Go
//
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - výpis typů hodnot konstant
// - výpis typů proměnných, do kterých je přiřazena konstanta

package main

import "fmt"

const (
	zprava1        = "foo"
	zprava2 string = "bar"
)

func main() {
	// v deklaraci těchto proměnných není explicitně zapsán typ
	// (ten je odvozen překladačem automaticky)
	var v1 = zprava1
	var v2 = zprava2

	// vypsat hodnoty proměnných
	fmt.Println(v1, v2)

	// vypsat typy konstant
	fmt.Printf("%T %T\n", zprava1, zprava2)

	// vypsat typy proměnných
	fmt.Printf("%T %T\n", v1, v2)
}
