// Transpiler GopherJS
//
// - implementace HTTP serveru, který dokáže na portu 8080 poskytnout
//   všechny soubory umístěné v aktuálním adresáři
// - souborům .js bude automaticky přidělen korektní MIME typ

package main

import (
	"fmt"
	"net/http"
)

func main() {
	const address = ":8080"

	const directory = http.Dir(".")

	fmt.Println("Starting HTTP server on address", address)
	err := http.ListenAndServe(address, http.FileServer(directory))

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
