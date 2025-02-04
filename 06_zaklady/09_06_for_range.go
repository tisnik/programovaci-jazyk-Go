// Základy programovacího jazyka Go
//
// - programová smyčka založená na klíčovém slovu range

package main

import "fmt"

func main() {
	// počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		fmt.Println("Counting:", i)
	}
}
