// Gorutiny a kanály v jazyce Go
//
// - deadlock při práci s kanály
// - čekání na prvek z kanálu, do kterého žádná gorutina neprovádí zápis

package main

import (
	"fmt"
)

func main() {
	// konstrukce kanálu
	channel := make(chan int)

	// čteme z kanálu, do kterého žádná gorutina neprovádí zápis
	code, status := <-channel
	fmt.Printf("received code: %t and status: %t\n", code, status)
}
