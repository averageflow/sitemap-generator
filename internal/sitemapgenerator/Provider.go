package sitemapgenerator

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

const sitemapLocation = "./sitemap.xml"
const sitemapTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
%s
</urlset> 
`

func shouldSkipLink(link string) bool {
	if strings.HasPrefix(link, "https://") {
		return true
	}

	if strings.HasPrefix(link, "http://") {
		return true
	}

	if strings.HasPrefix(link, "tel:") {
		return true
	}

	if strings.HasPrefix(link, "mailto:") {
		return true
	}

	return false
}

func CrawlWebsite(config *ApplicationConfig) string {
	var urls []*url.URL

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if shouldSkipLink(link) {
			return
		}

		known, err := c.HasVisited(link)
		if known || err != nil {
			return
		}

		err = e.Request.Visit(link)
		if err != nil {
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		urls = append(urls, r.URL)
		fmt.Println("Visiting: ", r.URL)
	})

	if config.IsMultilingual {
		err := c.Visit(fmt.Sprintf("%s/%s", config.Location, "en"))
		if err != nil {
			log.Println(err.Error())
		}

		err = c.Visit(fmt.Sprintf("%s/%s", config.Location, "nl"))
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		err := c.Visit(config.Location)
		if err != nil {
			log.Println(err.Error())
		}
	}

	// computation halts until all links have been traversed
	// in both locales
	c.Wait()

	var sitemapData string

	for i := range urls {
		sitemapData += fmt.Sprintf("<url><loc>%s</loc></url>\n\t", urls[i])
	}

	return fmt.Sprintf(sitemapTemplate, sitemapData)
}

func ReplaceSitemapFileWithNewData(data string) error {
	err := os.Remove(sitemapLocation)
	if err != nil {
		log.Println(err.Error())
	}

	err = ioutil.WriteFile(sitemapLocation, []byte(data), 0777)
	if err != nil {
		return err
	}

	return nil
}

func ParseConfigFlags() *ApplicationConfig {
	location := flag.String("location", "", "URL to initiate the crawling")
	isMultilingual := flag.Bool("multilingual", false, "should the crawler visit both /en and /nl variants")

	flag.Parse()

	if location == nil || isMultilingual == nil {
		flag.Usage()
		return nil
	}

	if *location == "" {
		flag.Usage()
		return nil
	}

	return &ApplicationConfig{
		Location:       *location,
		IsMultilingual: *isMultilingual,
	}
}
