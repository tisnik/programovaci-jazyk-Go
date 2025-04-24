// Jednotkové testy v jazyce Go
//
// - balíček s funkcí, která se bude testovat

package factorial

// testovaná funkce
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
