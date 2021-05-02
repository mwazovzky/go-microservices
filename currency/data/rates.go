package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

const ECB_URL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type ExchangeRates struct {
	logger hclog.Logger
	rates  map[string]float64
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

func NewExchangeRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{logger: l, rates: map[string]float64{}}

	err := er.getRates()

	return er, err
}

func (er *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := er.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currrency %s", base)
	}

	dr, ok := er.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currrency %s", dest)
	}

	return dr / br, nil
}

func (er *ExchangeRates) getRates() error {
	res, err := http.DefaultClient.Get(ECB_URL)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status code 200 got %d", res.StatusCode)
	}
	defer res.Body.Close()

	// parse xml data
	cd := Cubes{}
	xml.NewDecoder(res.Body).Decode(&cd)
	for _, c := range cd.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		er.rates[c.Currency] = r
	}

	er.rates["EUR"] = 1

	return nil
}
