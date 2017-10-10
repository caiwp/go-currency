package gocurrency

import "testing"

func TestCurrencyRate(t *testing.T) {
	f, e := CurrencyRate("USD", "CNY")
	if e != nil {
		t.Fatal(e)
		return
	}

	t.Error(f)
}

func TestConvent(t *testing.T) {
	buf := []byte(`"USDCNY=X",6.5731,"10/10/2017","10:57am"`)
	f, e := convent(buf)
	if e != nil {
		t.Fatal(e)
		return
	}
	t.Error(f)
}
