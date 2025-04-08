// Mapy v programovacím jazyku Go
//
// - průchod všemi prvky mapy

package main

import "fmt"

func main() {
	m1 := make(map[string]int)

	m1["nula"] = 0
	m1["jedna"] = 1
	m1["dva"] = 2
	m1["tri"] = 3
	m1["ctyri"] = 4
	m1["pet"] = 5
	m1["sest"] = 6

	// průchod všemi prvky mapy
	for key, value := range m1 {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
