// Ukazatele v jazyku Go
//
// - ukazatele na prvky záznamu

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	// proměnná typu záznam
	var u User

	u.id = 1
	u.name = "Pepek"
	u.surname = "Vyskoč"

	// tisk původní hodnoty záznamu
	fmt.Println(u)

	// ukazatel na jeden prvek záznamu
	var pName *string
	pName = &u.name

	// tisk ukazatele na prvek i hodnoty prvku
	fmt.Println(pName)
	fmt.Println(*pName)

	// modifikace prvku nepřímo přes ukazatel
	*pName = "Zdeněk"
	fmt.Println(*pName)
}
