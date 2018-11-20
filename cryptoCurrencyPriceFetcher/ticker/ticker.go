package main

import (
	"reflect"
	"encoding/json"
	"fmt"
	"strings"
)

const jsonStream = `{
    "data": {
        "278": {
            "id": 278,
            "name": "Quebecoin",
            "symbol": "QBC",
            "website_slug": "quebecoin",
            "rank": 1583,
            "circulating_supply": 15588776.0,
            "total_supply": 36375376.0,
            "max_supply": 42000000.0,
            "quotes": {
                "USD": {
                    "price": 0.0030329289,
                    "volume_24h": 15.1646447495,
                    "market_cap": 47280.0,
                    "percent_change_1h": -0.04,
                    "percent_change_24h": -1.31,
                    "percent_change_7d": -19.58
                },
                "AUD": {
                    "price": 0.0041354381,
                    "volume_24h": 20.6771902564,
                    "market_cap": 64466.0,
                    "percent_change_1h": -0.04,
                    "percent_change_24h": -1.31,
                    "percent_change_7d": -19.58
                }
            },
            "last_updated": 1542438994
        },
        "290": {
            "id": 290,
            "name": "BlueCoin",
            "symbol": "BLU",
            "website_slug": "bluecoin",
            "rank": 1284,
            "circulating_supply": 574683675.0,
            "total_supply": 574683675.0,
            "max_supply": null,
            "quotes": {
                "USD": {
                    "price": 0.0013786041,
                    "volume_24h": 1.7232550852,
                    "market_cap": 792261.0,
                    "percent_change_1h": -0.04,
                    "percent_change_24h": -1.31,
                    "percent_change_7d": -15.18
                },
                "AUD": {
                    "price": 0.0018797446,
                    "volume_24h": 2.3496807109,
                    "market_cap": 1080259.0,
                    "percent_change_1h": -0.04,
                    "percent_change_24h": -1.31,
                    "percent_change_7d": -15.18
                }
            },
            "last_updated": 1542438997
        }
    },
    "metadata": {
        "timestamp": 1542438264,
        "num_cryptocurrencies": 2081,
        "error": null
    }
}`

type lvl1 struct {
	Data     interface{} `json:"data"`
	Metadata lvl2M       `json:"metadata"`
}

type lvl2D struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	WebsiteSlug       string `json:"website_slug"`
	Rank              int    `json:"rank"`
	CirculatingSupply int    `json:"circulating_supply"`
	TotalSupply       int    `json:"total_supply"`
	MaxSupply         int    `json:"max_supply"`
	// Quotes            Info   `json:"quotes"`
	// "USD":
	//     "price": 0.0030329289,
	//     "volume_24h": 15.1646447495,
	//     "market_cap": 47280.0,
	//     "percent_change_1h": -0.04,
	//     "percent_change_24h": -1.31,
	//     "percent_change_7d": -19.5
	//  "AUD":
	//     "price": 0.0041354381,
	//     "volume_24h": 20.6771902564,
	//     "market_cap": 64466.0,
	//     "percent_change_1h": -0.04,
	//     "percent_change_24h": -1.31,
	//     "percent_change_7d": -19.5
	LastUpdated int `json:"last_updated"`
}

type lvl2M struct {
	Timestamp           int  `json:"timestamp"`
	NumCryptocurrencies int  `json:"num_cryptocurrencies"`
	Error               bool `json:"error"`
}

func main() {

	init := json.NewDecoder(strings.NewReader(jsonStream))

	// url := "https://api.coinmarketcap.com/v2/ticker/?limit=3"
	// resp, _ := http.Get(url)
	// init := json.NewDecoder(resp.Body)
	// defer resp.Body.Close()

	var c lvl1
	if err := init.Decode(&c); err != nil {
		panic(err)
	}
	next := c.Data
	fmt.Println("-------------------------")
	fmt.Println(reflect.TypeOf(next))
	fmt.Println("-------------------------")

    passInterface(next)
	newNext := next.(map[string]interface{})
	fmt.Println(reflect.TypeOf(newNext["278"]))
	fmt.Println("-------------------------")
}

func passInterface(v interface{}) {
    b, ok := v.(*[]byte)
    fmt.Println(ok)
    fmt.Println(b)
}