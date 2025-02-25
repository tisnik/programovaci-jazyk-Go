// Interpretry programovacího jazyka Go
//
// Interpret Gore
//
// - zdrojový soubor vygenerovaný interpretrem Gore
// - stav po deklaraci trojice proměnnych
// - zdrojový kód je zobrazen přesně tak, jak ho tiskne pseudopříkaz :print

package main

import "github.com/k0kubun/pp/v3"

func __gore_p(xs ...any) {
	for _, x := range xs {
		pp.Println(x)
	}
}
func main() { x := 42; y := 10; z := "foobar" }
