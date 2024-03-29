package main

import (
	"encoding/xml"
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
	// check sitemap url
	if len(os.Args) < 3 {
		log.Fatal("ERROR: please input sitemap URL and output file name as an argument.")
	}
	url := os.Args[1]
	sitemap := GetSitemap(url)
	var pages = make([]byte, 0, 100)
	for _, page := range sitemap {
		pages = append(pages, page.Loc...)
		pages = append(pages, "\n"...)
	}
	fileName := os.Args[2]
	OutputFile(pages, fileName)
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

func OutputFile(pages []byte, fileName string) {
	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()
	file.Write((pages))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
