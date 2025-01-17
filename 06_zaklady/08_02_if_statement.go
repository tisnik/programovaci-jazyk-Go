// Základy programovacího jazyka Go
//
// - rozhodovací konstrukce s lokální proměnnou

package main

import "fmt"

func main() {
	if j := 10; j < 11 {
		fmt.Println(j, "< 11")
	}
}
