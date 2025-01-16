// Základní datové typy v jazyce Go
//
// - pravdivostní hodnota deklarovaná jako lokální konstanta

package main

import "fmt"

func main() {
	// deklarace lokální konstanty
	const AutomaticAnswer = true

	// tisk hodnoty konstanty
	fmt.Println(AutomaticAnswer)
}
