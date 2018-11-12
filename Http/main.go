package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.Get returns 2 things. resp and err,
	// we don't need error in this case so we igore it with ", _"
	resp, _ := http.Get("https://www.w3schools.com/xml/simple.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
