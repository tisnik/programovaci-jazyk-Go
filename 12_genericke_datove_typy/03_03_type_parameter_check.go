// Generické datové typy v Go
//
// - vynucení konkrétního typového parametru
//   při volání funkce printValue
// - tento program nelze přeložit!

package main

import "fmt"

// funkce s typovým parametrem
func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue[int]("Programovací jazyk Go")
	printValue[[]string]('*')
	printValue[string](42)
	printValue[int](3.14)
	printValue[byte](1 + 2i)
	printValue[[]byte]([]int{1, 2, 3})
}
