// Jednotkové testy v jazyce Go
//
// - balíček s funkcí, která se bude testovat

package main

import "fmt"

// testovaná funkce
func add(x int32, y int32) int32 {
	return x - y
}

func main() {
	fmt.Println(add(1, 2))
}
