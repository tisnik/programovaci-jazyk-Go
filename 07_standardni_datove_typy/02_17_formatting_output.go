// Tisk celočíselných numerických hodnot s jejich naformátováním.

package main

import "fmt"

func main() {
	// proměnné různých celočíselných typů obsahujících záporné hodnoty
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

	var r1 rune = 'a'
	var r2 rune = '\x40'
	var r3 rune = '\n'
	var r4 rune = '\u03BB'

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

	// tisk numerické hodnoty formou znaku
	fmt.Println("\n%c")
	fmt.Printf("%c\n", r1)
	fmt.Printf("%c\n", r2)
	fmt.Printf("%c\n", r3)
	fmt.Printf("%c\n", r4)
}
