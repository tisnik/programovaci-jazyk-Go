// Generické datové typy v Go
//
// - generická varianta funkce pro porovnání
//   dvou hodnot stejného (generického) typu

package main

import "fmt"

type comparable interface {
	int | float64 | string
}

func compare[T comparable](x T, y T) bool {
	return x < y
}

func main() {
	fmt.Println(compare(1, 2))
	fmt.Println(compare(1.5, 2.6))
	fmt.Println(compare("foo", "bar"))
}
