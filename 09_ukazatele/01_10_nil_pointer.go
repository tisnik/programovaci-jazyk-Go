// Ukazatele v jazyku Go
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

// funkce, která modifikuje obsah proměnné přes ukazatel
func increment(ptr *int) {
	// nepřímá změna obsahu proměnné
	*ptr++
}

func main() {
	// deklarace ukazatele, který bude nastavený na nil
	var p *int

	// pokus o zvýšení hodnoty uložené na adrese, kterou
	// obsahuje ukazatel
	increment(p)
}
