// Generické datové typy v Go
//
// - funkce printValue akceptující
//   hodnotu libovolného typu
// - využití prázdného rozhraní pro typ parametru

package main

import "fmt"

// funkce akceptující hodnotu libovolného typu
func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {
	printValue("Programovací jazyk Go")
	printValue('*')
	printValue(42)
	printValue(3.14)
	printValue(1 + 2i)
	printValue([]int{1, 2, 3})
}
