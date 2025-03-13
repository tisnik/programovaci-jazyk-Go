// Technologie GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci CalcSum
// - kontrola typu argumentů předaných funkci CalcSum
// - provedení konverze na nativní typy jazyka Go
// - převedení výsledku zpět na JavaScriptový typ

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func CalcSum(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 2 {
		fmt.Printf("incorrect number of arguments %d, but just two are accepted\n", len(args))
		return nil
	}

	// kontrola typu předaných argumentů
	// (pozor - původní syntaxe zápisu smyčky)
	for index := 0; index < 2; index++ {
		typ := args[index].Type()
		if typ != js.TypeNumber {
			fmt.Printf("Argument #%d has incorrect type %s\n", index, typ.String())
			return nil
		}
	}

	// počet i typ argumentů je korektní
	// lze tedy provést jejich konverzi
	x := args[0].Int()
	y := args[1].Int()

	// vypočítat výsledek
	z := x + y

	// vrátit výsledek s explicitní konverzí
	// na typ kompatibilní s JavaScriptem
	return js.ValueOf(z)
}

func main() {
	fmt.Println("started")

	// export funkce CalcSum tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("calcSum", js.FuncOf(CalcSum))

	fmt.Println("finished")
}
