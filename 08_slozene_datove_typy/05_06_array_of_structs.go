// Struktury (záznamy) v programovacím jazyku Go
//
// - pole, jehož prvky jsou typu struktura (záznam)
// - přímá inicializace prvků pole

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// pole záznamů s inicializací jeho prvků
	var users = [3]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}

	// tisk obsahu celého pole
	fmt.Println(users)

	// druhé pole, překladač si odvodí jeho velikost
	var users2 = [...]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}

	// tisk obsahu celého pole
	fmt.Println(users2)
}
