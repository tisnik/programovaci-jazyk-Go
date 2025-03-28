// Ukazatele
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

func increment(ptr *int) {
	*ptr++
}

func main() {
	var p *int

	increment(p)
}
