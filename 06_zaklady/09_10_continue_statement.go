// Základy programovacího jazyka Go
//
// - příkaz continue v programových smyčkách

package main

import "fmt"

func main() {
	// počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		// přeskočit sudé hodnoty
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
