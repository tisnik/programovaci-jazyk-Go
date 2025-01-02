// Základní datové typy v jazyce Go
//
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk jmen typů proměnných i jejich (nulových) hodnot

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

	// tisk typů i hodnot všech proměnných
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", flags, flags)
	fmt.Printf("%T\t%v\n", c1, c1)
	fmt.Printf("%T\t%v\n", c2, c2)
	fmt.Printf("%T\t%v\n", e, e)
	fmt.Printf("%T\t%v\n", f1, f1)
	fmt.Printf("%T\t%v\n", f2, f2)
	fmt.Printf("%T\t%v\n", i1, i1)
	fmt.Printf("%T\t%v\n", i2, i2)
	fmt.Printf("%T\t%v\n", i3, i3)
	fmt.Printf("%T\t%v\n", i4, i4)
	fmt.Printf("%T\t%v\n", i5, i5)
	fmt.Printf("%T\t%v\n", znak, znak)
	fmt.Printf("%T\t%v\n", zprava, zprava)
	fmt.Printf("%T\t%v\n", u1, u1)
	fmt.Printf("%T\t%v\n", u2, u2)
	fmt.Printf("%T\t%v\n", u3, u3)
	fmt.Printf("%T\t%v\n", u4, u4)
	fmt.Printf("%T\t%v\n", u5, u5)
	fmt.Printf("%T\t%v\n", pointer, pointer)
}
