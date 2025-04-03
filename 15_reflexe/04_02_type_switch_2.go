// Reflexe v programovacím jazyku Go
//
// - konverze hodnoty předané v parametru typu any
//   na vybrané konkrétní datové typy
// - otestování, že konvertované hodnoty mají skutečně
//   očekávaný datový typ a lze s nimi provádět operace
//   specifické pro tento typ

package main

import "fmt"

func test_type(value any) {
	// konverze hodnoty předané v parametru typu "any"
	// na vybrané konkrétní datové typy int, bool a string
	// nakonec se provedou různé operace s konvertovanými hodnotami
	switch v := value.(type) {
	case int:
		fmt.Println("Integer value:", v*10)
	case bool:
		fmt.Println("Boolean value:", !v)
	case string:
		fmt.Println("String value:", "***"+v+"***")
	default:
		fmt.Println("Unsupported value")
	}
}

func main() {
	x := 42
	test_type(x)

	y := true
	test_type(y)

	z := "foobar"
	test_type(z)

	w := 1 + 2i
	test_type(w)
}
