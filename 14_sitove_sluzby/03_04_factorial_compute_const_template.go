// Vývoj síťových aplikací v programovacím jazyku Go
//
// - neměnná šablona stránky

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type FactorialPageDynContent struct {
	N      int64
	Result int64
}

func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

func getNFromForm(request *http.Request) (int64, error) {
	err := request.ParseForm()
	if err != nil {
		return 0, err
	}

	n, err := strconv.ParseInt(request.FormValue("n"), 10, 64)
	if err != nil {
		n = 0
	}
	return n, nil
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request, t *template.Template) {
	n, err := getNFromForm(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result := Factorial(n)
	dynData := FactorialPageDynContent{N: n, Result: result}

	err = t.Execute(writer, dynData)
	if err != nil {
		fmt.Println("Error executing template")
	}
}

func main() {
	// načtení šablony
	t, err := template.ParseFiles("factorial_compute.html")
	if err != nil {
		fmt.Println("Cannot load and parse template")
		os.Exit(1)
	}

	// registrace callback funkce volané při přístupu na endpoint
	http.HandleFunc("/", func(
		writer http.ResponseWriter, request *http.Request) {
		mainEndpoint(writer, request, t)
	})

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
