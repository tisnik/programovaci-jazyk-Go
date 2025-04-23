// Jednotkové testy v jazyce Go
//
// - balíček s funkcí, která se bude testovat

package main

import "fmt"

// testovaná funkce
func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(1, 2))
}
