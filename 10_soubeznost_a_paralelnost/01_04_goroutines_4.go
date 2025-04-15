// Gorutiny a kanály v jazyce Go
//
// - vytvoření a zavolání různých gorutin
// - gorutiny se volají z jiných gorutin
// - gorutiny budou zapisovat znaky na stejný textový řádek

package main

import (
	"fmt"
	"time"
)

// tisk sekvence znaků na textový řádek se zpožděním
func printChars() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
		time.Sleep(200 * time.Millisecond)
	}
}

// tisk sekvence teček na textový řádek se zpožděním
func printDots() {
	for i := 0; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
}

// tisk sekvence mezer na textový řádek se zpožděním
func printSpaces() {
	// vytvoření a zavolání gorutiny
	go printChars()

	// vytvoření a zavolání gorutiny
	go printDots()
	for i := 0; i < 60; i++ {
		fmt.Print(" ")
		time.Sleep(110 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main begin")

	// vytvoření a zavolání gorutiny
	go printSpaces()

	// naivní čekání na dokončení všech tří gorutin
	time.Sleep(6 * time.Second)

	fmt.Println("main end")
}
