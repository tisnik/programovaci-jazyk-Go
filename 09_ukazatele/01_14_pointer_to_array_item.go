// Ukazatele v jazyku Go
//
// - ukazatel na prvek pole

package main

import "fmt"

func main() {
	// deklarace a inicializace pole
	mesice := [12]string{
		"Leden",
		"Únor",
		"Březen",
		"Duben",
		"Květen",
		"Červen",
		"Červenec",
		"Srpen",
		"Září",
		"Říjen",
		"Listopad",
		"Prosinec"}

	// tisk původního obsahu pole
	fmt.Println(mesice)

	// deklarace ukazatele na zvolený prvek pole
	pThirdMonth := &mesice[2]

	// nepřímá modifikace prvku přes ukazatel
	*pThirdMonth = "March"

	// tisk nového obsahu pole
	fmt.Println(mesice)
}
