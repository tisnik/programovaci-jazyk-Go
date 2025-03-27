// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - HTTP server se specifikací MIME typu odpovědí

package main

import (
	"fmt"
	"net/http"
)

func endpointHTML(writer http.ResponseWriter, request *http.Request) {
	// nastavení hlavičky
	writer.Header().Set("Content-Type", "text/html")
	response := "<body><h1>Hello world!</h1></body>\n"
	// posílat budeme sekvenci bajtů
	writer.Write([]byte(response))
}

func endpointText(writer http.ResponseWriter, request *http.Request) {
	// nastavení hlavičky
	writer.Header().Set("Content-Type", "text/plain")
	response := "Hello world!\n"
	// posílat budeme sekvenci bajtů
	writer.Write([]byte(response))
}

func main() {
	// registrace callback funkce volané při přístupu na endpointy
	http.HandleFunc("/html", endpointHTML)
	http.HandleFunc("/text", endpointText)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
