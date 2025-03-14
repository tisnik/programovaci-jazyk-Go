// Základy programovacího jazyka Go
//
// - nekonečná smyčka realizovaná příkazem goto
//
// POZOR! nepoužívat, nejedná se o idiomatický kód

package main

import "fmt"

func main() {
opak:
	fmt.Println("Diamonds are forever")
	goto opak
}
