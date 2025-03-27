// Vývoj síťových aplikací v programovacím jazyku Go
//
// - druhá varianta webové aplikace pro výpočet faktoriálu

package main

import (
	"fmt"
	"html/template"
	"net/http"
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

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := strconv.ParseInt(request.FormValue("n"), 10, 64)
	if err != nil {
		n = 0
	}

	t, err := template.ParseFiles("factorial_compute.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := FactorialPageDynContent{N: n, Result: Factorial(n)}
	err = t.Execute(writer, dynData)
	if err != nil {
		fmt.Println("Error executing template")
	}
}

func main() {
	// registrace callback funkce volané při přístupu na endpoint
	http.HandleFunc("/", mainEndpoint)

	// spuštění HTTP serveru
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":8000", nil)
}
