// Hodnota nil v jazyce Go
//
// - proměnné různých typů, které jsou automaticky
//   nastaveny na hodnotu nil

package main

import "fmt"

func main() {
	// ukazatel
	var p *int

	// řez
	var s []int

	// mapa
	var m map[string]int

	// kanál
	var c chan int

	// rozhraní
	var a any

	fmt.Println(p, p == nil)
	fmt.Println(s, s == nil)
	fmt.Println(m, m == nil)
	fmt.Println(c, c == nil)
	fmt.Println(a, a == nil)
}
