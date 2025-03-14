// Základy programovacího jazyka Go
//
// - nekonečná smyčka realizovaná příkazem goto
// - explicitní výskok ze smyčky dalším příkazem goto
//
// POZOR! nepoužívat, nejedná se o idiomatický kód

package main

import "fmt"

func main() {
	i := 0
opak:
	fmt.Println("Diamonds are forever")
	i += 1
	if i > 10 {
		goto konec
	}
	goto opak
konec:
}
