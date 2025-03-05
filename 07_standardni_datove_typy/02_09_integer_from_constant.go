// Základní datové typy v jazyce Go
//
// - lokální proměnné typu celé číslo bez znaménka
// - inicializace lokálních proměnných globální konstantou

package main

import "fmt"

// globální celočíselná konstanta
const Answer = 42

func main() {
	// typ odpovídající osmibitovému, 16bitovému, 32bitovému
	// a 64bitovému číslu bez znaménka
	var a uint8 = Answer
	var b uint16 = Answer
	var c uint32 = Answer
	var d uint64 = Answer

	// zjištění, jaký výchozí formát je použit pro tisk celočíselných
	// hodnot se znaménkem
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
