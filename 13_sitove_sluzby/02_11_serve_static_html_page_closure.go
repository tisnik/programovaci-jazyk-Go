// Vývoj síťových aplikací v programovacím jazyku Go
//
// - standardní balíček net/http
// - server, který vrací statické HTML stránky načtené ze souborů
// - realizace přes uzávěr

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

func staticPage(filename string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		sendStaticPage(writer, filename)
	}
}

func main() {
	// registrace callback funkce volané při přístupu na endpointy
	http.HandleFunc("/", staticPage("index.html"))
	http.HandleFunc("/missing", staticPage("missing.html"))

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
