// Celočíselné konstanty zapisované ve dvojkové, osmičkové a šestnáctkové
// (hexadecimální) soustavě. Použití znaku _ pro vizuální oddělení číslic.

package main

import "fmt"

func main() {
	// oddělovač u hodnoty zapsané v desítkové soustavě
	var a uint32 = 1_000
	var b uint32 = 1_000_0000

	// pozor: nelze použít
	// var c uint32 = _1_000_0000
	// var d uint32 = 1_000_0000_

	fmt.Println(a)
	fmt.Println(b)

	// binární, osmičková a hexadecimální soustava
	var e uint8 = 0b0101_0101
	var f uint8 = 0_10
	var g uint8 = 0o1_0
	var h uint32 = 0xFFFF_FFFF

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
