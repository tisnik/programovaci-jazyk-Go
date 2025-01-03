// Celočíselné konstanty zapisované ve dvojkové, osmičkové a šestnáctkové
// (hexadecimální) soustavě.

package main

import "fmt"

func main() {
	// binární, osmičková a hexadecimální soustava
	var a uint8 = 0b0101
	var b uint8 = 010
	var c uint8 = 0o10
	var d uint8 = 0x10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// použít lze i záporné hodnoty
	var e int8 = -0b0101
	var f int8 = -010
	var g int8 = -0o10
	var h int8 = -0x10

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
