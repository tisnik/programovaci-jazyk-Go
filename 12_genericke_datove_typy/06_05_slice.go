// Generické datové typy v Go
//
// - generické funkce car a cdr pro řezy

package main

import "fmt"

// car vrátí první prvek řezu
func car[T any](s []T) T {
	return s[0]
}

// cdr vrátí řez bez prvního prvku
func cdr[T any](s []T) []T {
	return s[1:]
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println(car(s))
	fmt.Println(cdr(s))
}
