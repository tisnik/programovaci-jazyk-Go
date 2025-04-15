// Gorutiny a kanály v jazyce Go
//
// - vytvoření a zavolání dvou gorutin

package main

import "fmt"

// funkce, která bude zavolána v gorutině
func message(id int) {
	fmt.Printf("gorutina %d\n", id)
}

func main() {
	fmt.Println("main begin")

	// vytvoření a zavolání gorutiny
	go message(1)

	// vytvoření a zavolání gorutiny
	go message(2)

	fmt.Println("main end")
}
