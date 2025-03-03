// První kroky v jazyku Go
//
// - program typu "Hello, world!" ve variantě naprogramované
//   v jazyku Go.

// specifikace jména balíčku (zde musíme použít "main")
package main

// import konstant, typů a funkcí z jiného balíčku
import "fmt"

// deklarace funkce nazvané main (vstupní bod do programu)
func main() {
	// volání funkce náležející do jiného balíčku
	// s využitím takzvané tečkové notace
	fmt.Println("Hello, world!")
}
