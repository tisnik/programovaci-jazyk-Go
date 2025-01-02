// Základní datové typy v jazyce Go
//
// - využití celočíselných datových typů se znaménkem

package main

import "fmt"

func main() {
	// typ odpovídající osmibitovému, 16bitovému, 32bitovému
	// a 64bitovému číslu se znaménkem
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int64 = -1000000

	// typ "rune" je taktéž celočíselný typ se znaménkem
	var r1 rune = 'a'
	var r2 rune = '\x40'
	var r3 rune = '\n'
	var r4 rune = '\u03BB'

	// zjištění, jaký výchozí formát je použit pro tisk celočíselných
	// hodnot se znaménkem
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// vytištění run - z pohledu Go se jedná o celočíselné hodnoty, nikoli
	// o znaky
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}
