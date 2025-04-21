// Gorutiny a kanály v jazyce Go
//
// - deadlock při práci s kanály
// - zápis do kanálu, ze kterého žádná gorutina neprovádí čtení

package main

func main() {
	// konstrukce kanálu
	channel := make(chan int)

	// zápis do kanálu, ze kterého žádná gorutina neprovádí čtení
	channel <- 42
}
