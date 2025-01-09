// Základy programovacího jazyka Go
//
// - základní vlastnosti numerických konstant

package main

import "fmt"

const (
	// numerické konstanty
	x = 10
	y = 3.1415927
	z = 2 + 0i
)

func main() {
	// přiřazení numerických konstant do proměnných
	// různých typů
	var v1 int = x
	var v2 float32 = y
	var v3 float64 = y
	var v4 complex64 = y
	var v5 complex128 = y
	var v6 int = z
	var v7 float32 = z
	var v8 float64 = z
	var v9 complex64 = z
	var v10 complex128 = z

	// výpis hodnot všech deseti proměnných
	fmt.Println("v1=", v1)
	fmt.Println("v2=", v2)
	fmt.Println("v3=", v3)
	fmt.Println("v4=", v4)
	fmt.Println("v5=", v5)
	fmt.Println("v6=", v6)
	fmt.Println("v7=", v7)
	fmt.Println("v8=", v8)
	fmt.Println("v9=", v9)
	fmt.Println("v10=", v10)
}
