// Technologie WebAssembly
//
// - program typu "Hello, world!" ve variantě naprogramované
//   v jazyku Go.
// - importuje se standardní balíček "fmt" a volá se funkce fmt.Println
// - tento program lze přeložit do nativního spustitelného
//   souboru i do WebAssembly

// specifikace jména balíčku (zde musíme použít "main")
package main

// import konstant, typů a funkcí z jiného balíčku
import "fmt"

// deklarace funkce nazvané main
func main() {
	// volání funkce náležející do jiného balíčku
	// s využitím takzvané tečkové notace
	fmt.Println("Hello, world!")
}
