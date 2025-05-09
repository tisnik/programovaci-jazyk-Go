// Generické datové typy v Go
//
// - generická varianta funkce pro součet dvou
//   hodnot stejného (generického) typu
// - odvození typů od typů základních
// - této funkci předáváme hodnoty odvozených typů
// - aproximace datových typů

package main

import "fmt"

type numeric interface {
	// aproximace datových typů
	~int | ~float64 | ~complex128
}

// generická varianta funkce pro součet dvou
// hodnot stejného (generického) typu
func add[T numeric](x T, y T) T {
	return x + y
}

// odvozené datové typy
type myInt int

type myFloat float64

type myComplex complex128

func main() {
	// proměnné získané z odvozených datových typů
	var x myInt = 42
	var y myFloat = 3.14
	var z myComplex = 1 + 2i

	fmt.Println(add(x, x))
	fmt.Println(add(y, y))
	fmt.Println(add(z, z))
}
