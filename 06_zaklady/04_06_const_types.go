// Základy programovacího jazyka Go
//
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - přiřazení pojmenovaných konstant do proměnných
// - použití pojmenovaných konstant ve výrazu

package main

import "fmt"

const (
	zprava1        = "foo"
	zprava2 string = "bar"
)

func main() {
	var v1, v2 string
	// přiřazení pojmenovaných konstant do proměnných
	v1 = zprava1
	v2 = zprava2
	fmt.Println(v1, v2)

	var v3 string

	// pojmenované konstanty použité ve výrazu
	v3 = zprava1 + zprava2 + "."
	fmt.Println(v3)
}
