// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch s vyhodnocovaným výrazem typu string
// - několik konstant zapsaných ve stejné větvi case

package main

import "fmt"

func command(x string) string {
	switch x {
	case "":
		return "chybí příkaz"
	case "help", "info":
		return "nápověda"
	case "bye", "exit", "quit":
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
