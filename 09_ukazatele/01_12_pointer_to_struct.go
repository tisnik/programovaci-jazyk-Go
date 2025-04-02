// Ukazatele v jazyku Go
//
// - ukazatel na záznam (strukturu)
// - změna prvku záznamu nepřímo přes ukazatel

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

	// ukazatel na záznam s jeho inicializací
	var pUser *User
	pUser = &u

	// tisk ukazatele na záznam i samotného záznamu
	fmt.Println(pUser)
	fmt.Println(*pUser)

	// modifikace prvku záznamu nepřímo přes ukazatel
	(*pUser).id = 10000
	fmt.Println(*pUser)

	// modifikace prvku záznamu nepřímo přes ukazatel
	pUser.id = 20000
	fmt.Println(*pUser)
}
