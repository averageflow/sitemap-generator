package main

import (
	"log"

	"github.com/averageflow/sitemap-generator.git/internal/sitemapgenerator"
)

func main() {
	config := sitemapgenerator.ParseConfigFlags()
	if config == nil {
		return
	}

	sitemapData := sitemapgenerator.CrawlWebsite(config)

	err := sitemapgenerator.ReplaceSitemapFileWithNewData(sitemapData)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
