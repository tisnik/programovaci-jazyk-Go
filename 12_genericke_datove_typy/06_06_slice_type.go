// Generické datové typy v Go
//
// - definice generického datového typu Slice
// - generické funkce car a cdr pro řezy

package main

import "fmt"

type Slice[T any] []T

// car vrátí první prvek řezu
func car[T any](s Slice[T]) T {
	return s[0]
}

// cdr vrátí řez bez prvního prvku
func cdr[T any](s Slice[T]) []T {
	return s[1:]
}

func main() {
	s := Slice[int]{1, 2, 3}

	fmt.Println(s)
	fmt.Println(car(s))
	fmt.Println(cdr(s))
}
