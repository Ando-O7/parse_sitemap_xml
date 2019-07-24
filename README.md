# parse_sitemap_xml
Extract URL list of Sitemap to file.

## Usage
```
$ git clone git@github.com:Ando-O7/parse_sitemap_xml.git
$ cd parse_sitemap_xml
$ go build main.go
$ ./main [URL of sitemap] [Output file name]
```
## Example
sitemap URL : https://www.google.com/sitemap.xml  
file name : google_sitemap.txt

```
$ ./main https://www.google.com/sitemap.xml google_sitemap.txt
$ less google_sitemap.txt
https://www.google.com/gmail/sitemap.xml
https://www.google.com/forms/sitemaps.xml
https://www.google.com/slides/sitemaps.xml
etc...
```

## Note
The type needs to be changed by the structure of sitemap.

### case 1

For example, Google has the following structure: it is using a sitemapindex file.
```
<sitemapindex xmlns="http://www.google.com/schemas/sitemap/0.84">
    <sitemap>
        <loc>https://www.google.com/gmail/sitemap.xml</loc>
    </sitemap>
    <sitemap>
        <loc>https://www.google.com/forms/sitemaps.xml</loc>
    </sitemap>
    <sitemap>
        <loc>https://www.google.com/slides/sitemaps.xml</loc>
    </sitemap>
    etc...
```
Type is as follows.
```
type Sitemap struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Pages   []Page   `xml:"sitemap"`
}
```

### case 2

In the case of other site, it has the following structure. default sitemap.
```
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>http://example.org/</loc>
    </url>
    <url>
        <loc>http://example.org/xxxxx/</loc>
    </url>
    <url>
        <loc>http://example.org/yyyyy</loc>
    </url>
    etc...
```
Type is as follows.
```
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Pages   []Page   `xml:"url"`
}
```

## TODO
- Correspond to various structures of sitemap.