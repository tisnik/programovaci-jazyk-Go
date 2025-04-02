// Ukazatele v jazyku Go
//
// - ukazatel na jiný ukazatel

package main

import "fmt"

func main() {
	// proměnná typu celé číslo
	var i int

	// ukazatel na celé číslo
	var pi *int

	// ukazatel na ukazatel na celé číslo
	var ppi **int

	// tisk původních (nulových) hodnot proměnné i ukazatelů
	fmt.Printf("i: (%T) = %v\n", i, i)
	fmt.Printf("pi: (%T) = %v\n", pi, pi)
	fmt.Printf("ppi: (%T) = %v\n", ppi, ppi)

	fmt.Println()

	// inicializace proměnné i obou ukazatelů
	i = 42
	pi = &i
	ppi = &pi

	// tisk nových (nenulových) hodnot proměnné i ukazatelů
	fmt.Printf("i: (%T) = %v\n", i, i)
	fmt.Printf("pi: (%T) = %v\n", pi, pi)
	fmt.Printf("ppi: (%T) = %v\n", ppi, ppi)

	fmt.Println()

	// přístup k proměnné přímo, nepřímo přes ukazatel a nepřímo
	// přes dva ukazatele
	fmt.Printf("    i = %v\n", i)
	fmt.Printf("  *pi = %v\n", *pi)
	fmt.Printf("**ppi = %v\n", **ppi)
}
