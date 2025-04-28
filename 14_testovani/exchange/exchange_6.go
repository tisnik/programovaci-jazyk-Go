// metoda pro převod měn (již ne funkce)
func (g *ExchangeGetter) exchange(amount float64, code string) float64 {
	rate := g.get_exchange_rate(code)
	return rate * amount
}
