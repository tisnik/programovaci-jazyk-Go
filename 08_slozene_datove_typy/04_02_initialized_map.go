// Mapy v programovacím jazyku Go
//
// - deklarace prázdné mapy
// - přidání dvojic klíč-hodnota do mapy

package main

import "fmt"

func main() {
	var m1 map[string]int = make(map[string]int)
	fmt.Println(m1)

	m1["nula"] = 0
	m1["jedna"] = 1
	m1["dva"] = 2
	m1["tri"] = 3
	m1["ctyri"] = 4
	m1["pet"] = 5
	m1["sest"] = 6

	fmt.Println(m1)
}
