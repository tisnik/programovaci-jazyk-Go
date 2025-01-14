// Tisk celočíselných numerických hodnot s jejich naformátováním.

package main

import "fmt"

func main() {
	// proměnné různých celočíselných typů obsahujících kladné hodnoty
	var a uint8 = 20
	var b uint16 = 2000
	var c uint32 = 20000
	var d uint32 = 2000000

	// specifikace dekadického formátu
	fmt.Println("%d")
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

	// specifikace dekadického formátu s určením počtu znaků
	fmt.Println("\n%5d")
	fmt.Printf("%5d\n", a)
	fmt.Printf("%5d\n", b)
	fmt.Printf("%5d\n", c)
	fmt.Printf("%5d\n", d)

	// specifikace dekadického formátu s určením počtu znaků
	// s vynucením tisku nul na začátku hodnot s menším počtem znaků
	fmt.Println("\n%05d")
	fmt.Printf("%05d\n", a)
	fmt.Printf("%05d\n", b)
	fmt.Printf("%05d\n", c)
	fmt.Printf("%05d\n", d)

	// specifikace dekadického formátu s určením počtu znaků
	// a zarovnáním doleva
	fmt.Println("\n%-5d")
	fmt.Printf("%-5d\n", a)
	fmt.Printf("%-5d\n", b)
	fmt.Printf("%-5d\n", c)
	fmt.Printf("%-5d\n", d)

	// specifikace dekadického formátu s určením počtu znaků
	// a vynucením tisku znaménka
	fmt.Println("\n%+5d")
	fmt.Printf("%+5d\n", a)
	fmt.Printf("%+5d\n", b)
	fmt.Printf("%+5d\n", c)
	fmt.Printf("%+5d\n", d)

	// specifikace hexadecimálního formátu, přičemž hexadecimální
	// číslice jsou tisknuty malými znaky
	fmt.Println("\n%x")
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	fmt.Printf("%x\n", c)
	fmt.Printf("%x\n", d)

	// specifikace hexadecimálního formátu, přičemž hexadecimální
	// číslice jsou tisknuty velkými znaky
	fmt.Println("\n%X")
	fmt.Printf("%X\n", a)
	fmt.Printf("%X\n", b)
	fmt.Printf("%X\n", c)
	fmt.Printf("%X\n", d)

	// specifikace osmičkovéh formátu
	fmt.Println("\n%o")
	fmt.Printf("%o\n", a)
	fmt.Printf("%o\n", b)
	fmt.Printf("%o\n", c)
	fmt.Printf("%o\n", d)

	// specifikace binárního formátu
	fmt.Println("\n%b")
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", d)
}
