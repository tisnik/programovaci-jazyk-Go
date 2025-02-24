// Řetězce v programovacím jazyku Go
//
// - řetězec s Unicode znaky
// - průchod jednotlivými znaky řetězce
// - tisk těchto znaků

package main

import "fmt"

func main() {
	var s string = "Příliš žluťoučký kůň"

	for _, c := range s {
		fmt.Printf("%c", c)
	}
}
