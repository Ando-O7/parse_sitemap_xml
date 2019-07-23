package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Loc string `xml:"loc"`
}
type Sitemap struct {
	// Change according to sitemap structure
	XMLName xml.Name `xml:"sitemapindex"`
	Pages   []Page   `xml:"sitemap"`
}

func main() {
	url := "https://www.google.com/sitemap.xml"
	sitemap := GetSitemap(url)
	for _, page := range sitemap {
		fmt.Println(page.Loc)
	}
}

func GetSitemap(url string) []Page {
	response, err := http.Get(url)
	checkErr(err)

	XMLdata, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	var sitemap Sitemap
	xml.Unmarshal(XMLdata, &sitemap)
	return sitemap.Pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
