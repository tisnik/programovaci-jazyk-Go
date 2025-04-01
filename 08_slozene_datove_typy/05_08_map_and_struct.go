// Struktury (záznamy) v programovacím jazyku Go
//
// - mapa, která obsahuje uložené záznamy
// - jako klíče jsou použity taktéž záznamy

package main

import "fmt"

// Key je uživatelsky definovaná datová struktura
type Key struct {
	id   uint32
	role string
}

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// inicializace prázdné mapy
	m1 := make(map[Key]User)

	// tisk prázdné mapy
	fmt.Println(m1)

	// přidání prvku do mapy
	m1[Key{1, "admin"}] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	// přidání prvku do mapy
	m1[Key{2, "user"}] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	// tisk naplněné mapy
	fmt.Println(m1)
}
