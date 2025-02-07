// Interpretry programovacího jazyka Go
//
// Interpret Gore
//
// - zdrojový soubor vygenerovaný interpretrem Gore
// - stav po zadání příkazu :import fmt
// - zdrojový kód je zobrazen přesně tak, jak ho tiskne pseudopříkaz :print

package main

import (
    "github.com/k0kubun/pp/v3"
    "fmt"
)

func __gore_p(xs ...any) {
    for _, x := range xs {
        pp.Println(x)
    }
}
func main() {}
