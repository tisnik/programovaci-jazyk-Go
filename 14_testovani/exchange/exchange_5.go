// konstruktor
func NewExchangeGetter(g ExchangeDataGetter) *ExchangeGetter {
	return &ExchangeGetter{get_exchange_rate: g}
}
