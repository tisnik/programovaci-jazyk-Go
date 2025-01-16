// Základní datové typy v jazyce Go
//
// - získání reálné a imaginární složky komplexního čísla
// - varianta programu pro typ complex128

package main

import "fmt"

func main() {
	// zpracovávané komplexní číslo
	var a complex128 = -1.5 + 3.4i

	// tisk komplexního čísla i jeho složek
	fmt.Println("a=", a)
	fmt.Println("real(a)=", real(a))
	fmt.Println("imag(a)=", imag(a))

	fmt.Println()

	// tisk typů komplexního čísla i jeho složek
	fmt.Printf("type(a):       %T\n", a)
	fmt.Printf("type(real(a)): %T\n", real(a))
	fmt.Printf("type(imag(a)): %T\n", imag(a))
}
