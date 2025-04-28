// funkce pro převod měn
func exchange(amount float64, code string) float64 {
	rate := get_exchange_rate(code)
	return rate * amount
}
