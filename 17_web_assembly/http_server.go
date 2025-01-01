// Technologie WebAssembly
//
// - implementace HTTP serveru, který dokáže na portu 8080 poskytnout
//   všechny soubory umístěné v aktuálním adresáři
// - souborům WASM bude automaticky přidělen korektní MIME typ

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// adresa a port, na kterém budou dostupné soubory poskytované HTTP
	// serverem
	const address = ":8080"

	// adresář se soubory, které má poskytovat HTTP server
	const directory = http.Dir(".")

	fmt.Println("Starting HTTP server on address", address)
	// spuštění HTTP serveru na portu 8080
	err := http.ListenAndServe(address, http.FileServer(directory))

	// zpracování chyby při spuštění serveru a/nebo jeho běhu
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
