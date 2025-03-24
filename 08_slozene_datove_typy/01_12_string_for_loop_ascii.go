// Řetězce v programovacím jazyku Go
//
// - řetězec obsahující jen ASCII znaky
// - průchod jednotlivými bajty tvořícími řetězec

package main

import "fmt"

func main() {
	var s string = "Hello world!"

	// průchod přes jednotlivé bajty
	for i := 0; i < len(s); i++ {
		// tisk bajtu tak, jakoby se jednalo
		// o kód ASCII znaku
		fmt.Printf("%c", s[i])
	}
}
