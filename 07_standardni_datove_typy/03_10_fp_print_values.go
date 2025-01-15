// Základní datové typy v jazyce Go
//
// - tisk numerických hodnot s plovoucí řádovou čárkou
//   s jejich naformátováním.

package main

import "fmt"

func main() {
	// hodnota, kterou budeme vypisovat
	var f float64 = 1.23

	// naformátování s využitím řádové tečky,
	// ale nikoli v exponenciálním tvaru
	fmt.Println("Formát %f (bez exponentu)")
	fmt.Printf("%f\n", f)
	fmt.Printf("%10f\n", f)
	fmt.Printf("%11.7f\n", f)
	fmt.Printf("%5.0f\n", f)
	fmt.Printf("%.0f\n", f)

	fmt.Println()

	// naformátování v exponenciálním tvaru
	fmt.Println("Formát %e (s exponentem)")
	fmt.Printf("%e\n", f)
	fmt.Printf("%10e\n", f)
	fmt.Printf("%11.7e\n", f)
	fmt.Printf("%5.0e\n", f)
	fmt.Printf("%.0e\n", f)

	fmt.Println()

	// naformátování s využitím řádové tečky nebo v exponenciálním tvaru
	// (podle velikosti konkrétní hodnoty)
	fmt.Println("Formát %g (hodnoty blízko nuly)")
	fmt.Printf("%g\n", f)
	fmt.Printf("%10g\n", f)
	fmt.Printf("%11.7g\n", f)
	fmt.Printf("%5.0g\n", f)
	fmt.Printf("%.0g\n", f)

	fmt.Println()
	f = 10e7

	// naformátování s využitím řádové tečky nebo v exponenciálním tvaru
	// (nyní s vysokou hodnotou)
	fmt.Println("Formát %g (velké hodnoty)")
	fmt.Printf("%g\n", f)
	fmt.Printf("%10g\n", f)
	fmt.Printf("%11.7g\n", f)
	fmt.Printf("%6.0g\n", f)
	fmt.Printf("%.0g\n", f)

	fmt.Println()
	f = 10e-7

	// naformátování s využitím řádové tečky nebo v exponenciálním tvaru
	// (nyní s hodnotou blízkou nule)
	fmt.Println("Formát %g (malé hodnoty)")
	fmt.Printf("%g\n", f)
	fmt.Printf("%10g\n", f)
	fmt.Printf("%11.7g\n", f)
	fmt.Printf("%6.0g\n", f)
	fmt.Printf("%.0g\n", f)
}
