// Základy programovacího jazyka Go
//
// - příklad rozvětvení pomocí konstrukce switch
// - deklarace lokální proměnné přímo v konstrukci switch

package main

import "fmt"

func main() {
	switch i := 5; i {
	case 4:
		fmt.Println(i, "= 4")
	case 5:
		fmt.Println(i, "= 5")
	default:
		fmt.Println(i, "jiná hodnota než 4 nebo 5")
	}
}
