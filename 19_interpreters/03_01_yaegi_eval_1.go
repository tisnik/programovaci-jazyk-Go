// Interpretry programovacího jazyka Go
//
// Interpret Yaegi
//
// - využití interpretru Yaegi
// - vyhodnocení jednořádkového výrazu

package main

import (
	"fmt"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// inicializace nového interpretru a vyhodnocení předaného výrazu
func RunYaegi(script string) any {
	// inicializace nového interpretru
	interpreter := interp.New(interp.Options{})

	interpreter.Use(stdlib.Symbols)

	// vyhodnocení předaného výrazu
	ret, err := interpreter.Eval(script)
	if err != nil {
		panic(err)
	}

	return ret
}

func main() {
	fmt.Println(RunYaegi("6*7"))
}
