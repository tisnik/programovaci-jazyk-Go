// Datový typ bool a pravdivostní hodnoty v programovacím jazyku Go.
// Překladač jazyka Go neumožní automatickou konverzi mezi nějakým
// datovým typem odlišným od bool na pravdivostní hodnotu.

package main

import "fmt"

func main() {
	// korektní deklarace a inicializace proměnné
	var a bool = true
	var b bool = false

	// nekorektní deklarace a inicializace proměnné
	var c bool = 0
	var d bool = ""

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
