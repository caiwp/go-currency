package gocurrency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/text/currency"
)

const baseURL = "https://openexchangerates.org/api/latest.json?app_id=%s"
const appID = ""

func CurrencyRate(from, to string) (float64, error) {
	if _, err := currency.ParseISO(from); err != nil {
		return 0, err
	}

	if _, err := currency.ParseISO(to); err != nil {
		return 0, err
	}

	url := fmt.Sprintf(baseURL, appID)
	buf, err := request(url)
	if err != nil {
		return 0, err
	}

	rates, err := unmarshal(buf)
	if err != nil {
		return 0, err
	}

	fi, ok := rates[from]
	if !ok {
		return 0, fmt.Errorf("currency rate not found %v", from)
	}
	ti, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("currency rate not found %v", to)
	}

	fv, ok := fi.(float64)
	if !ok {
		return 0, fmt.Errorf("currency rate is not float64 %T %v", fi, fi)
	}
	tv, ok := ti.(float64)
	if !ok {
		return 0, fmt.Errorf("currency rate is not float64 %T %v", ti, ti)
	}

	return compute(tv, fv)
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

func unmarshal(buf []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(buf, &data); err != nil {
		return nil, err
	}
	rates, ok := data["rates"]
	if !ok {
		return nil, fmt.Errorf("rates not found")
	}
	v, y := rates.(map[string]interface{})
	if !y {
		return nil, fmt.Errorf("rates not right")
	}
	return v, nil
}

func compute(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("value is 0")
	}
	v := x / y
	s := fmt.Sprintf("%.3f", v)
	return strconv.ParseFloat(s, 64)
}
