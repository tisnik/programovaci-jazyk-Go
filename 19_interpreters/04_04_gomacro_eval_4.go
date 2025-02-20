// Interpretry programovacího jazyka Go
//
// Interpret Gomacro
//
// - využití interpretru Gomacro
// - vyhodnocení víceřadkového skriptu
// - skript obsahuje deklaraci funkce, která je volána

package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

// skript, který se bude interpretovat
const script = `
func multiply(a, b int) int {
    return a*b
}

x:=6
y:=7
multiply(x, y)
`

// inicializace nového interpretru a vyhodnocení předaného skriptu
func RunGomacro(script string) any {
	// inicializace nového interpretru
	interp := fast.New()

	// vyhodnocení předaného skriptu
	vals, _ := interp.Eval(script)
	if len(vals) < 1 {
		return "no value"
	}
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro(script))
}
