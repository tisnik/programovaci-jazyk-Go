// Základní datové typy v jazyce Go
//
// - komplexní hodnota deklarovaná jako lokální konstanta
// - tisk této hodnoty

package main

import "fmt"

func main() {
	// deklarace lokální konstanty
	const c = 1 + 0.5i

	// tisk hodnoty konstanty
	fmt.Println(c)
}
