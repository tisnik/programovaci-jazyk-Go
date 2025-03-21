// Mapy v programovacím jazyku Go
//
// - zkrácená deklarace prázdné mapy
// - mapa má předalokovanou kapacitu
// - přidání dvojic klíč-hodnota do mapy

package main

import "fmt"

func main() {
	const capacity = 100
	m1 := make(map[string]int, capacity)
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
