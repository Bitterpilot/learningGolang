package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Data struct which contains an array of data
type Data struct {
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

func findCoinByID(ID int, s Data) (int, string, string, string, error) {
	for i := 0; i < len(s.Coins); i++ {
		if s.Coins[i].ID == ID {
			return s.Coins[i].ID,
				s.Coins[i].Symbol,
				s.Coins[i].Name,
				s.Coins[i].WebsiteSlug,
				nil
		}
	}
	return -1, "", "", "", errors.New("Can't Find Coin ID")
}

type indexPage struct {
	Title string
	List  []Coin
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.coinmarketcap.com/v2/listings/")
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var coins Data
	json.Unmarshal(byteValue, &coins)

	for i := 0; i < 5; i++ {
		d, s, n, w, err := findCoinByID(i, coins)
		if err != nil {
			fmt.Println("No details For this ID")
		} else {
			fmt.Println("details:", d, s, n, w)
		}
	}

	page := indexPage{"List of Crypto Currencies", coins.Coins}
	t, err := template.ParseFiles("templates/index.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
