// Interpretry programovacího jazyka Go
//
// Interpret Gore
//
// - zdrojový soubor vygenerovaný interpretrem Gore
// - povolení automatických importů
// - stav po zadání příkazu fmt.Println()
// - zdrojový kód je zobrazen přesně tak, jak ho tiskne pseudopříkaz :print

package main

import (
    "fmt"

    "github.com/k0kubun/pp/v3"
)

func __gore_p(xs ...any) {
    for _, x := range xs {
        pp.Println(x)
    }
}
func main() { _, _ = fmt.Println("Hello") }
