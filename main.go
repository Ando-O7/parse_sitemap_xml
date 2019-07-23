package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Loc string `xml:"loc"`
}
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Pages   []Page   `xml:"url"`
}

func main() {
	// Read xml file
	url := "https://www.google.com/sitemap.xml"
	response, err := http.Get(url)

	// check error
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", response.Status)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
