// Generické datové typy v Go
//
// - otestování, že oba parametry funkce
//   stejného generického typu jsou shodného
//   aktuálního typu
// - tento příklad nelze přeložit!

package main

import "fmt"

type additive interface {
	int | string
}

func add[T additive](x T, y T) T {
	return x + y
}

func main() {
	// volání funkce s předáním hodnot různých typů
	fmt.Println(add(1, 2))
	fmt.Println(add(1, "foo"))
	fmt.Println(add("bar", 1))
	fmt.Println(add("foo", "bar"))
}
