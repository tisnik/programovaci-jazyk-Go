// Základní datové typy v jazyce Go
//
// - hodnota s plovoucí čárkou deklarovaná jako lokální konstanta
// - explicitní převod na typ float32
// - tisk typu této konstanty

package main

import "fmt"

func main() {
	// deklarace lokální konstanty
	const Pi = float32(3.1415)

	// tisk typu konstanty
	fmt.Printf("%T\n", Pi)
}
