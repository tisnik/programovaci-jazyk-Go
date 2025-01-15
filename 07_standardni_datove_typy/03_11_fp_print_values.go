// Základní datové typy v jazyce Go
//
// - tisk numerických hodnot s plovoucí řádovou čárkou
//   s jejich naformátováním.

package main

import "fmt"

func main() {
	// hodnota, kterou budeme vypisovat
	var f float64 = 1.25

	// naformátování s využitím binárního exponentu
	fmt.Println("Formát %b (binární mantisa i exponent)")
	fmt.Printf("%b\n", f)
	fmt.Printf("%10b\n", f)
	fmt.Printf("%11.7b\n", f)
	fmt.Printf("%5.0b\n", f)
	fmt.Printf("%.0b\n", f)

	fmt.Println()

	// naformátování s využitím hexadecimální mantisy
	// a binárního exponentu
	fmt.Println("Formát %x (hexadecimální mantisa a binární exponent)")
	fmt.Printf("%x\n", f)
	fmt.Printf("%10x\n", f)
	fmt.Printf("%11.7x\n", f)
	fmt.Printf("%5.0x\n", f)
	fmt.Printf("%.0x\n", f)
}
