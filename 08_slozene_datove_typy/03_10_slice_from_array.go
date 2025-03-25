// Řezy v programovacím jazyku Go
//
// - vytvoření řezů z pole

package main

import "fmt"

func main() {
	// pole se sto prvky
	var a1 [100]byte

	// pole se sto prvky
	var a2 [100]int32

	fmt.Printf("Délka pole 1:   %d\n", len(a1))
	fmt.Printf("Délka pole 2:   %d\n", len(a2))
	fmt.Println()

	// řez vytvořený z prvního pole, který obsahuje
	// pohled na všechny prvky pole
	var slice0 []byte = a1[:]
	fmt.Printf("Délka řezu 0:     %d\n", len(slice0))
	fmt.Printf("Kapacita řezu 0:  %d\n", cap(slice0))
	fmt.Println()

	// řez vytvořený z prvního pole, který obsahuje
	// pohled na deset prvků s indexy od 0 do 9
	var slice1 []byte = a1[:10]
	fmt.Printf("Délka řezu 1:     %d\n", len(slice1))
	fmt.Printf("Kapacita řezu 1:  %d\n", cap(slice1))
	fmt.Println()

	// řez vytvořený z prvního pole, který obsahuje
	// pohled na deset prvků s indexy od 10 do 19
	var slice2 []byte = a1[10:20]
	fmt.Printf("Délka řezu 2:     %d\n", len(slice2))
	fmt.Printf("Kapacita řezu 2:  %d\n", cap(slice2))
	fmt.Println()

	// řez vytvořený z prvního pole, který obsahuje
	// pohled na jiných deset prvků s indexy od 20 do 29
	var slice3 = a1[20:30]
	fmt.Printf("Délka řezu 3:     %d\n", len(slice3))
	fmt.Printf("Kapacita řezu 3:  %d\n", cap(slice3))
	fmt.Println()

	// řez vytvořený z prvního pole, který obsahuje
	// pohled na dalších deset prvků s indexy od 30 do 39
	slice4 := a1[30:40]
	fmt.Printf("Délka řezu 4:     %d\n", len(slice4))
	fmt.Printf("Kapacita řezu 4:  %d\n", cap(slice4))
}
