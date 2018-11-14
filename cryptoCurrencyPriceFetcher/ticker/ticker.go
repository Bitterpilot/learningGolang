package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Data struct which contains an array of data
type Data struct {
	Coins []Coin `json:"data"`
}

// Coin struct which contains the id, name, symbol
// and slug of a coin
type Coin struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	WebsiteSlug string      `json:"website_slug"`
	Rank        int         `json:"rank"`
	CirSupply   float64     `json:"circulating_supply"`
	AvalSupply  float64     `json:"total_supply"`
	MaxSupply   float64     `json:"max_supply"`
	Quotes      interface{} `json:"quotes"`
	LastUpdate  int64       `json:"last_updated"`
}

type quote struct {
	USD CurrencyGeneric `json:"USD"`
	AUD CurrencyGeneric `json:"AUD"`
}

// CurrencyGeneric ...
type CurrencyGeneric struct {
	Price        float64 `json:"price"`
	Volume24     float64 `json:"volume_24h"`
	Cap          float64 `json:"market_cap"`
	PctChange1h  float64 `json:"percent_change_1h"`
	PctChange24h float64 `json:"percent_change_24h"`
	PctChange7d  float64 `json:"percent_change_7d"`
}

func currancyParser(in interface{}) {
	iface := in
	v := reflect.ValueOf(iface)
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Println(key.Interface())
		fmt.Println(val.Interface())

		valTyp := reflect.TypeOf(val.Interface()).Kind()
		if valTyp != reflect.Float64 {
			currancyParser(val.Interface())
		}
	}

}

func main() {
	// https://api.coinmarketcap.com/v2/ticker/
	// options
	// preface ?
	// seperator &
	//
	// start int
	// limit int
	// sort string
	// structure string
	// convert string

	url := "https://api.coinmarketcap.com/v2/ticker/?convert=AUD&start=101&limit=10&sort=id&structure=array"
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var coins Data
	json.Unmarshal(byteValue, &coins)

	// fmt.Println(coins.Coins[1])

	// get time from last update
	// tm := time.Unix(coins.Coins[1].LastUpdate, 0)
	// fmt.Println(tm)

	currancyParser(coins.Coins[0].Quotes)
}
