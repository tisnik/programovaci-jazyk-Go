// Základní datové typy v jazyce Go
//
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk jmen typů proměnných

package main

import "fmt"

func main() {
	// lokální proměnné automaticky inicializované na nulové hodnoty
	var (
		a       any
		b       bool
		flags   byte
		c1      complex64
		c2      complex128
		e       error
		f1      float32
		f2      float64
		i1      int
		i2      int8
		i3      int16
		i4      int32
		i5      int64
		znak    rune
		zprava  string
		u1      uint
		u2      uint8
		u3      uint16
		u4      uint32
		u5      uint64
		pointer uintptr
	)

	// tisk typů všech proměnných
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", flags)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", i3)
	fmt.Printf("%T\n", i4)
	fmt.Printf("%T\n", i5)
	fmt.Printf("%T\n", znak)
	fmt.Printf("%T\n", zprava)
	fmt.Printf("%T\n", u1)
	fmt.Printf("%T\n", u2)
	fmt.Printf("%T\n", u3)
	fmt.Printf("%T\n", u4)
	fmt.Printf("%T\n", u5)
	fmt.Printf("%T\n", pointer)
}
