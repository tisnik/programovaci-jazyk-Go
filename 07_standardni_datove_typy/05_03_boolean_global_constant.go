// Základní datové typy v jazyce Go
//
// - pravdivostní hodnota deklarovaná jako globální konstanta

package main

import "fmt"

// deklarace globální konstanty
const AutomaticAnswer = true

func main() {
	// tisk hodnoty konstanty
	fmt.Println(AutomaticAnswer)
}
