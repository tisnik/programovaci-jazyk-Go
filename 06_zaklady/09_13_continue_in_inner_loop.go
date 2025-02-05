// Základy programovacího jazyka Go
//
// - vnořené programové smyčky
// - příkaz continue pro pokračování další iteraci vnitřní

package main

import "fmt"

func main() {
	// vnější počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		// vnitřní počítaná smyčka pro hodnoty i od 0 do 9
		for j := range 10 {
			result := (i + 1) * (j + 1)
			fmt.Printf("%3d ", result)
			if result == 42 {
				fmt.Print("odpověď nalezena!")
				// přeskok zbytku vnitřní smyčky
				continue
			}
		}
		// odřádkování
		fmt.Println()
	}
}
