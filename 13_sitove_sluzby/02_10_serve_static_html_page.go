// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - server, který vrací statické HTML stránky načtené ze souborů

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendStaticPage(writer http.ResponseWriter, filename string) {
	body, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	sendStaticPage(writer, "index.html")
}

func missingPageEndpoint(writer http.ResponseWriter, request *http.Request) {
	sendStaticPage(writer, "missing.html")
}

func main() {
	// registrace callback funkce volané při přístupu na endpointy
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/missing", missingPageEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
