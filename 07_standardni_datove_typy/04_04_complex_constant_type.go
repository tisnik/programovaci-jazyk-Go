// Základní datové typy v jazyce Go
//
// - komplexní hodnota deklarovaná jako lokální konstanta
// - tisk typu této konstanty

package main

import "fmt"

func main() {
	// deklarace lokální konstanty
	const c = 1 + 0.5i

	// tisk typu konstanty
	fmt.Printf("%T\n", c)
}
