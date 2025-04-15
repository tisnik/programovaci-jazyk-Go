// Generické datové typy v Go
//
// - funkce, která spojí textové reprezentace
//   všech prvků řezu
// - tento příklad nelze přeložit!

package main

import "fmt"

func join[T any](items ...T) (result string) {
	// průchod všemi prvky řezu
	for _, value := range items {
		// spojení textových reprezentací prvků
		result += value.String()
		result += ","
	}
	return result
}

func main() {
	fmt.Println(join("first", "second"))
}
