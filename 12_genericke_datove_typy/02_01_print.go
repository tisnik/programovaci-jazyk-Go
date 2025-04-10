// Generické datové typy v Go
//
// - printValue jakožto běžná negenerická funkce

package main

import "fmt"

// běžná negenerická funkce
func printValue(value string) {
	fmt.Println(value)
}

func main() {
	printValue("Učíme se jazyk Go!")
}
