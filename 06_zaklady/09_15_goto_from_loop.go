// Základy programovacího jazyka Go
//
// - jednoduchá programová smyčka
// - výskok ze smyčky pomocí goto

package main

import "fmt"

func main() {
	i := 0
	// nekonečná smyčka
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			//výskok ze smyčky
			goto Exit
		}
	}
	// návěští
Exit:
}
