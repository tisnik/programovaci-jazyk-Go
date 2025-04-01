// Struktury (záznamy) v programovacím jazyku Go
//
// - mapa, která obsahuje uložené záznamy
// - jako klíče jsou použity řetězce

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// inicializace prázdné mapy
	m1 := make(map[string]User)

	// tisk prázdné mapy
	fmt.Println(m1)

	// přidání prvku do mapy
	m1["prvni"] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	// přidání prvku do mapy
	m1["druhy"] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	// tisk naplněné mapy
	fmt.Println(m1)
}
