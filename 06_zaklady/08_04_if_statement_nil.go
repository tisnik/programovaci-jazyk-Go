// Základy programovacího jazyka Go
//
// - řídicí konstrukce if a hodnoty, které nejsou typu boolean

package main

func main() {
	if nil {
		println("true")
	}

	if !nil {
		println("false")
	}

	if "" {
		println("true")
	}

	if !"" {
		println("false")
	}

	var b1 bool = nil

	if b1 {
		println("true")
	}

	if !b1 {
		println("false")
	}
}
