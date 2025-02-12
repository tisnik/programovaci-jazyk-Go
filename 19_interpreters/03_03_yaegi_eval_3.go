// Interpretry programovacího jazyka Go
//
// Interpret Yaegi
//
// - využití interpretru Yaegi
// - vyhodnocení víceřadkového skriptu
// - skript vrací hodnotu

package main

import (
	"fmt"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// skript, který se bude interpretovat
const script = `
x:=6
y:=7
z:=x*y
z
`

// inicializace nového interpretru a vyhodnocení předaného skriptu
func RunYaegi(script string) any {
	// inicializace nového interpretru
	interpreter := interp.New(interp.Options{})

	interpreter.Use(stdlib.Symbols)

	// vyhodnocení předaného skriptu
	ret, err := interpreter.Eval(script)
	if err != nil {
		panic(err)
	}

	return ret
}

func main() {
	fmt.Println(RunYaegi(script))
}
