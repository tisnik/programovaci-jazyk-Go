// Základy programovacího jazyka Go
//
// - rozdíl mezi veřejnou (public) konstantou a konstantou soukromou (private)
// - zápis deklarace konstant do společného bloku

package main

import "fmt"

// jméno veřejné konstanty začíná velkým písmenem
// jméno soukromé konstanty začíná malým písmenem
const (
	VerejnaKonstanta  = "foo"
	soukromaKonstanta = "bar"
)

func main() {
	// zde pochopitelně můžeme použít obě konstanty - jsou viditelné
	fmt.Println(VerejnaKonstanta, soukromaKonstanta)
}
