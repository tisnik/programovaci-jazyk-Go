// Jednotkové testy v jazyce Go
//
// - balíček s funkcí, která se bude testovat

package main

import "fmt"

// testovaná funkce
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
