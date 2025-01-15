// Základní datové typy v jazyce Go
//
// - komplexní hodnota deklarovaná jako globální konstanta
// - tisk této hodnoty

package main

import "fmt"

// deklarace globální konstanty
const c = 1 + 0.5i

func main() {
	// tisk hodnoty konstanty
	fmt.Println(c)
}
