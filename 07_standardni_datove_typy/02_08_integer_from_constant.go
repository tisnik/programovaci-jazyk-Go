//

package main

import "fmt"

// globální celočíselná konstanta
const Answer = 42

func main() {
	// typ odpovídající osmibitovému, 16bitovému, 32bitovému
	// a 64bitovému číslu se znaménkem
	var a int8 = Answer
	var b int16 = Answer
	var c int32 = Answer
	var d int64 = Answer

	// typ "rune" je taktéž celočíselný typ se znaménkem
	var r1 rune = Answer

	// zjištění, jaký výchozí formát je použit pro tisk celočíselných
	// hodnot se znaménkem
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// vytištění run - z pohledu Go se jedná o celočíselné hodnoty, nikoli
	// o znaky
	fmt.Println(r1)
}
