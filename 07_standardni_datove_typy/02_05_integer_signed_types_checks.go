// Základní datové typy v jazyce Go
//
// - kontroly prováděné překladačem při přiřazování celočíselných konstant
//   do proměnných

package main

import "fmt"

func main() {
	// konstanta větší, než je osmibitový rozsah hodnot se znaménkem
	var a int8 = -1000

	// konstanta větší, než je 16bitový rozsah hodnot se znaménkem
	var b int16 = -100000

	// konstanta větší, než je 32bitový rozsah hodnot se znaménkem
	var c int32 = -10000000000

	// konstanta větší, než je 64bitový rozsah hodnot se znaménkem
	var d int32 = -10000000000000000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
