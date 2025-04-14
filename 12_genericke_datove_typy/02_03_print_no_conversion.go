// Generické datové typy v Go
//
// - Go nepodporuje automatické přetypování
//   na odvozený datový typ
// - tento příklad nelze přeložit!

package main

import "fmt"

// nový datový typ odvozený od řetězce
type Value string

// funkce akceptující hodnotu typu Value
// ovšem nikoli například řetězec
func printValue(value Value) {
	fmt.Println(value)
}

func main() {
	v := "Programovací jazyk Go" // string
	printValue(v)
}
