// Řetězce v programovacím jazyku Go
//
// - řetězec s Unicode znaky
// - průchod jednotlivými znaky řetězce
// - tisk kódů znaků

package main

import "fmt"

func main() {
	var s string = "Příliš žluťoučký kůň"

	// průchod řetězcem po runách
	for _, c := range s {
		fmt.Printf("%02x ", c)
	}
}
