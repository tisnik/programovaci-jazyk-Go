// Gorutiny a kanály v jazyce Go
//
// - použití kanálu pro čekání na dokončení gorutiny

package main

import "fmt"

// funkce, která bude zavolána v gorutině
func message(id int, channel chan int) {
	fmt.Printf("gorutina %d\n", id)

	// zápis libovolné hodnoty do kanálu
	channel <- 1
}

func main() {
	// konstrukce kanálu
	channel := make(chan int)

	fmt.Println("main begin")
	// vytvoření gorutiny, předání kanálu
	go message(1, channel)

	fmt.Println("waiting...")

	// blokující čtení z kanálu
	code, status := <-channel

	fmt.Printf("received code: %d and status: %t\n", code, status)
	fmt.Println("main end")
}
