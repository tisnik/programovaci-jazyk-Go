// Generické datové typy v Go
//
// - typový parametr u funkce printValue

package main

import "fmt"

// typový parametr u funkce printValue
func printValue[T any](value T) {
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
