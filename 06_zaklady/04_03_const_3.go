// Základy programovacího jazyka Go
//
// - rozdíl mezi veřejnou (public) konstantou a konstantou soukromou (private)

package main

import "fmt"

// jméno veřejné konstanty začíná velkým písmenem
const VerejnaKonstanta = "foo"

// jméno soukromé konstanty začíná malým písmenem
const soukromaKonstanta = "bar"

func main() {
	// zde pochopitelně můžeme použít obě konstanty - jsou viditelné
	fmt.Println(VerejnaKonstanta, soukromaKonstanta)
}
