package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3}
	s := []int{1, 2, 3}

	fmt.Printf("a type:  %T\n", a)
	fmt.Printf("s type:  %T\n", s)

	fmt.Printf("a value: %v\n", a)
	fmt.Printf("s value: %v\n", s)
}
