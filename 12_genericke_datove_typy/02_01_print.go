// Generické datové typy v Go
//
// - funkce printValue, která je deklarována
//   jakožto běžná negenerická funkce
// - zavolání této funkce

package main

import "fmt"

// běžná negenerická funkce
func printValue(value string) {
	fmt.Println(value)
}

func main() {
	printValue("Učíme se jazyk Go!")
}
