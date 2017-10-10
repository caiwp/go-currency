package gocurrency

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/text/currency"
)

//http://download.finance.yahoo.com/d/quotes.csv?e=.csv&f=sl1d1t1&s=USDCNY=X
const baseUrl = "http://download.finance.yahoo.com/d/quotes.csv?e=.csv&f=sl1d1t1&s=%s=X"

func CurrencyRate(from, to string) (float64, error) {
	if _, err := currency.ParseISO(from); err != nil {
		return 0, err
	}

	if _, err := currency.ParseISO(to); err != nil {
		return 0, err
	}

	url := fmt.Sprintf(baseUrl, from+to)

	buf, err := request(url)
	if err != nil {
		return 0, err
	}

	return convent(buf)
}

func request(url string) ([]byte, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func convent(buf []byte) (float64, error) {
	s := string(buf[11:17])
	return strconv.ParseFloat(s, 64)
}
