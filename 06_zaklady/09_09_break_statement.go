// Základy programovacího jazyka Go
//
// - příkaz break v programových smyčkách

package main

import "fmt"

func main() {
	// počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}
