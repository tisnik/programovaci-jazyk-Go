// Základy programovacího jazyka Go
//
// - deklarace proměnné "true", která přepíše původní hodnotu "true"
//
// POZOR: tento způsob není doporučený a je dobré vědět, kdy ho použít
//        a především, kdy ho naopak nepoužít.

package main

import "fmt"

func main() {
	// deklarace proměnné nazvané "true", tím v tomto balíčku ztratíme
	// přístup k původní hodnotě true (což většinou nechceme!)
	var true string = "pravda"

	// deklarace proměnné, která je ve skutečnosti typu "string"
	// (překladač nyní použije hodnotu proměnné "true")
	var answer string = true
	fmt.Println(answer)
}
