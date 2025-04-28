// různé způsoby převodu měny
// - na základě reálného kurzovního lístku
// - na základě obsahu souboru
// - pouze mockované chování

g := NewExchangeGetter(get_exchange_rate_from_url)
fmt.Printf("%5.2f\n", g.exchange(10, "USD"))

g2 := NewExchangeGetter(get_exchange_rate_from_file)
fmt.Printf("%5.2f\n", g2.exchange(10, "USD"))

g3 := NewExchangeGetter(mocked_get_exchange_rate)
fmt.Printf("%5.2f\n", g2.exchange(10, "USD"))
