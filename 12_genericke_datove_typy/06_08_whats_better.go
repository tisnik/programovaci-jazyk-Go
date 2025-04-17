// Generické datové typy v Go
//
// - rozdíl mezi generickou funkcí Print
//   a metodou Print

package main

import "fmt"

type Slice[T any] []T

// generická funkce Print
func (s Slice[T]) Print() {
	for _, value := range s {
		fmt.Println(value)
	}
}

// metoda Print
func Print[T any](s Slice[T]) {
	for _, value := range s {
		fmt.Println(value)
	}
}

func main() {
	s := Slice[int]{1, 2, 3}

	Print(s)

	fmt.Println()

	s.Print()
}
