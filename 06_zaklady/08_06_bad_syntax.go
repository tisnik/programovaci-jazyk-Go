// Základy programovacího jazyka Go
//
// - příklady špatné syntaxe konstrukce if-else.
// POZOR: tento příklad nebude možné přeložit

package main

func classify_char(c rune) string {
	if c >= 'a' && c <= 'z'
	{
		return "male pismeno"
	}
	else if c >= 'A' && c <= 'Z' {
		return "velke pismeno"
	}
	else
	{
		return "neco jineho"
	}
}

func main() {
	println(classify_char('a'))
	println(classify_char('Z'))
	println(classify_char('?'))
}
