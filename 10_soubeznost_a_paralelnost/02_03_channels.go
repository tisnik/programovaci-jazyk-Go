// Gorutiny a kanály v jazyce Go
//
// - vytvoření a zavolání tří gorutin
// - použití kanálu pro čekání na dokončení gorutiny

package main

import (
	"fmt"
	"time"
)

// funkce, která bude zavolána v gorutině
func message(id int, channel chan int) {
	fmt.Printf("gorutina %d\n", id)

	// počkáme přibližně dvě sekundy
	time.Sleep(2000 * time.Millisecond)

	// zápis libovolné hodnoty do kanálu_
	channel <- id
}

func main() {
	// konstrukce kanálu
	channel := make(chan int)

	fmt.Println("main begin")
	// vytvoření gorutiny, předání kanálu
	go message(1, channel)
	go message(2, channel)
	go message(3, channel)

	fmt.Println("waiting...")

	// blokující čtení z kanálu
	for range 3 {
		code, status := <-channel

		fmt.Printf("received code: %d and status: %t\n", code, status)
	}
	fmt.Println("main end")
}
