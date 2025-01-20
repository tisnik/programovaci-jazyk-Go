// Základy programovacího jazyka Go
//
// - příklad rozvětvení pomocí konstrukce switch

package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 4:
		fmt.Println(i, "= 4")
	case 5:
		fmt.Println(i, "= 5")
	default:
		fmt.Println(i, "jiná hodnota než 4 nebo 5")
	}
}
