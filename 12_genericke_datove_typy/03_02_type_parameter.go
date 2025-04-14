// Generické datové typy v Go
//
// - vynucení konkrétního typového parametru
//   při volání funkce printValue
// - vynucení konkrétního typu parametru funkce
//   při jejím volání

package main

import "fmt"

// funkce s typovým parametrem
func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	// vynucení konkrétního typu parametru funkce při jejím volání
	printValue[string]("Programovací jazyk Go")
	printValue[rune]('*')
	printValue[int](42)
	printValue[float32](3.14)
	printValue[complex64](1 + 2i)
	printValue[[]int]([]int{1, 2, 3})
}
