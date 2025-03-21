// Mapy v programovacím jazyku Go
//
// - smazání existujícího prvku z mapy
// - smazání neexistujícího prvku z mapy

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

	delete(m1, "dva")
	fmt.Println(m1)

	delete(m1, "neco jineho")
	fmt.Println(m1)
}
