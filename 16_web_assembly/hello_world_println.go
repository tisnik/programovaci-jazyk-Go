// Technologie WebAssembly
//
// - program typu "Hello, world!" ve variantě naprogramované
//   v jazyku Go.
// - používá se pouze funkce println a neimportuje se žádný
//   další balíček
// - tento program lze přeložit do nativního spustitelného
//   souboru i do WebAssembly

// specifikace jména balíčku (zde musíme použít "main")
package main

// deklarace funkce nazvané main
func main() {
	// volání funkce náležející do builtin balíčku
	// bez nutnosti importu balíčku fmt
	println("Hello, world!")
}
