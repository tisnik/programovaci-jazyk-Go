// Generické datové typy v Go
//
// - funkce, která spojí textové reprezentace
//   všech prvků řezu
// - realizace pro rozhraní TextLike

package main

import "fmt"

type TextLike interface {
	String() string
}

type Employee struct {
	name    string
	surname string
}

func (employee Employee) String() string {
	return employee.name + " " + employee.surname
}

func join[T TextLike](items ...T) (result string) {
	// průchod všemi prvky
	for _, value := range items {
		// převod na řetězec s připojením
		result += value.String()
		result += ","
	}
	return result
}

func main() {
	var e1 Employee = Employee{
		name:    "Pepek",
		surname: "Vyskoč",
	}

	var e2 Employee = Employee{
		name:    "Eda",
		surname: "Wasserfall",
	}

	fmt.Println(join(e1, e2))
}
