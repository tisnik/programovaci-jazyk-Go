// Základní datové typy v jazyce Go
//
// - pravdivostní hodnota deklarovaná jako lokální konstanta
// - tisk typu konstanty (ne její hodnoty)

package main

import "fmt"

func main() {
	// deklarace lokální konstanty
	const AutomaticAnswer = true

	// tisk typu konstanty
	fmt.Printf("%T\n", AutomaticAnswer)
}
