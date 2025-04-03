// Interpretry programovacího jazyka Go
//
// Interpret Gomacro
//
// - využití interpretru Gomacro
// - vyhodnocení jednořádkového výrazu

package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

// inicializace nového interpretru a vyhodnocení předaného výrazu
func RunGomacro(script string) any {
	// inicializace nového interpretru
	interp := fast.New()

	// vyhodnocení předaného výrazu
	vals, _ := interp.Eval(script)
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro("6*7"))
}
