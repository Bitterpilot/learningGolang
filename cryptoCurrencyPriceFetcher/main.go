package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Coins struct which contains an array of data
type Coins struct {
	Coins []Coin `json:"data"`
}

// Coin struct which contains the id, name, symbol
// and slug of a coin
type Coin struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	WebsiteSlug string `json:"website_slug"`
}

func findCoinByID(ID int, set Coins) (int, string, string, string, error) {
	for i := 0; i < len(set.Coins); i++ {
		if set.Coins[i].ID == ID {
			return set.Coins[i].ID,
				set.Coins[i].Symbol,
				set.Coins[i].Name,
				set.Coins[i].WebsiteSlug,
				nil
		}
	}
	return -1, "", "", "", errors.New("Can't Find Coin ID")
}

func main() {
	resp, err := http.Get("https://api.coinmarketcap.com/v2/listings/")
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var coins Coins
	json.Unmarshal(byteValue, &coins)

	for i := 0; i < 16; i++ {
		d, s, n, w, err := findCoinByID(i, coins)
		if err != nil {
			fmt.Println("No details For this ID")
		} else {
			fmt.Println("details:", d, s, n, w)
		}
	}
}
