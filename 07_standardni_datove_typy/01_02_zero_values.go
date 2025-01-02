// Základní datové typy v jazyce Go
//
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk výchozích (nulových) hodnot těchto proměnných

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

	fmt.Println("any       ", a)
	fmt.Println("bool      ", b)
	fmt.Println("byte      ", flags)
	fmt.Println("complex64 ", c1)
	fmt.Println("complex128", c2)
	fmt.Println("error     ", e)
	fmt.Println("float32   ", f1)
	fmt.Println("float64   ", f2)
	fmt.Println("int       ", i1)
	fmt.Println("int8      ", i2)
	fmt.Println("int16     ", i3)
	fmt.Println("int32     ", i4)
	fmt.Println("int64     ", i5)
	fmt.Println("rune      ", znak)
	fmt.Println("string    ", zprava)
	fmt.Println("uint      ", u1)
	fmt.Println("uint8     ", u2)
	fmt.Println("uint16    ", u3)
	fmt.Println("uint32    ", u4)
	fmt.Println("uint64    ", u5)
	fmt.Println("uintptr   ", pointer)
}
