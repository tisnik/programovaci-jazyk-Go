// Mapy v programovacím jazyku Go
//
// - přečtení prvku se známým klíčem
// - přečtení prvku s neznámým klíčem

package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m1["nula"] = 0
	m1["jedna"] = 1
	m1["dva"] = 2
	m1["tri"] = 3
	m1["ctyri"] = 4
	m1["pet"] = 5
	m1["sest"] = 6

	fmt.Println(m1)

	value1 := m1["dva"]
	fmt.Println(value1)

	value2 := m1["sto"]
	fmt.Println(value2)
}
