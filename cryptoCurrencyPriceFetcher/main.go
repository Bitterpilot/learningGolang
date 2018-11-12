package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Coins struct which contains an array of data
type Coins struct {
	Coins []Coin `json:"data"`
}

// Coin struct which contains the id, name, symbol
// and slug of a coin
type Coin struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	WebsiteSlug string `json:"website_slug"`
}

func main() {
	// jsonFile, _ := os.Open("sample.json")
	// defer jsonFile.Close()
	resp, err := http.Get("https://api.coinmarketcap.com/v2/listings/")
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()
	fmt.Println("Successfully Opened sample.json")

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var coins Coins

	json.Unmarshal(byteValue, &coins)

	for i := 0; i < len(coins.Coins); i++ {
		fmt.Println("ID: " + strconv.Itoa(coins.Coins[i].Id))
		fmt.Println("Symbol: " + coins.Coins[i].Symbol)
		fmt.Println("Name: " + coins.Coins[i].Name)
		fmt.Println("slug: " + coins.Coins[i].WebsiteSlug)
	}
}
