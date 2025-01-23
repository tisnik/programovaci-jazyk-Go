// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch s vyhodnocovaným výrazem typu string.

package main

import "fmt"

func command(x string) string {
	switch x {
	case "":
		return "chybí příkaz"
	case "help":
		fallthrough
	case "info":
		return "nápověda"
	case "bye":
		fallthrough
	case "exit":
		fallthrough
	case "quit":
		return "končím"
	default:
		return "neznámý příkaz"
	}
	return ""
}

func main() {
	fmt.Println(command(""))
	fmt.Println(command("bzz bzz bzz"))
	fmt.Println(command("bye"))
	fmt.Println(command("quit"))
	fmt.Println(command("exit"))
}
