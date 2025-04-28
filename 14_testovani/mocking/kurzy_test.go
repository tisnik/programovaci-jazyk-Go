// Testování v Go
//
// - otestování metody ExchangeGetter.exchange

package main

import (
	"testing"
)

func TestExchange(t *testing.T) {
	e := NewExchangeGetter(mocked_get_exchange_rate)
	expected := 215.0
	result := e.exchange(10.0, "USD")
	if result != expected {
		t.Errorf("Expected 0! == %f, but got %f instead", expected, result)
	}
}
