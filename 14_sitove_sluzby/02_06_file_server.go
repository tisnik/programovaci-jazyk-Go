// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - HTTP server vracející statický obsah

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// registrace callback funkce volané při přístupu k souborům
	http.Handle("/", http.FileServer(http.Dir("")))

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
