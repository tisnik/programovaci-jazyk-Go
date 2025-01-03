// Základní datové typy v jazyce Go
//
// - využití celočíselných datových typů bez znaménka

package main

import "fmt"

func main() {
	// typ odpovídající osmibitovému, 16bitovému, 32bitovému
	// a 64bitovému číslu bez znaménka
	var a uint8 = 10
	var b uint16 = 1000
	var c uint32 = 10000
	var d uint64 = 1000000

	// zjištění, jaký výchozí formát je použit pro tisk celočíselných
	// hodnot bez znaménka
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
