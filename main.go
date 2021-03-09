package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	//"github.com/gocolly/colly"
	"log"
	"time"
)
type news struct {
	Title string
	URL  string
	Contents string
	CrawledAt time.Time
}

func main() {
    fmt.Println("colly 爬虫开始!")
	c := colly.NewCollector(
		colly.AllowedDomains("www.baidu.com"),
		)

    fmt.Println(time.Now())
	log.Println("aa")
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("Visited", r.Request.Body)
		fmt.Println("Visited", r.Headers)
		fmt.Println("Visited", r.StatusCode)
	})
	c.OnHTML("body", func(element *colly.HTMLElement) {
		n := news{}
		n.Title = element.ChildText("div")
		log.Println(n)

	})
	c.Visit("https://www.baidu.com")

	resp, err := goquery.NewDocument("http://gold.3g.cnfol.com/")
	if err != nil {
		fmt.Println(err)
	}
	p := resp.Find("ul")
	s := p.Eq(6).Find("a")
	s.Each(func(i int, content *goquery.Selection) {
		a, _ := content.Attr("href")
		fmt.Println(a)
	})


}