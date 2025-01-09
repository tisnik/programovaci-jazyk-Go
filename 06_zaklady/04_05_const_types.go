// Základy programovacího jazyka Go
//
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty

package main

import "fmt"

const (
	zprava1        = "foo"
	zprava2 string = "bar"
)

func main() {
	// obě konstanty je možné vypsat
	// na standardní výstup
	fmt.Println(zprava1, zprava2)
}
