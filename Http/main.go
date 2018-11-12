package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type breakfastMenu struct {
	Name []name `xml:"food"`
}

type name struct {
	Nam string `xml:"name"`
}

func main() {
	// http.Get returns 2 things. resp and err,
	// we don't need error in this case so we igore it with ", _"
	resp, _ := http.Get("https://www.w3schools.com/xml/simple.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s breakfastMenu
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Name)
}
