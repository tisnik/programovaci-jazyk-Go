// Základy programovacího jazyka Go
//
// - základní varianta rozhodovací konstrukce "if"

package main

import "fmt"

func main() {
	var i int = 5

	if i < 6 {
		fmt.Println(i, "< 6")
	}
}
