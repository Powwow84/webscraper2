package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

type WeaponStats struct {
	SecHead string
	Guns string
	WeaponTypePopularity string
	
}

func main() {

	WeaponStats := WeaponStats{}

	scrapeUrl := "https://www.wzranked.com/wz2/meta/guns"

	collector := colly.NewCollector(colly.AllowedDomains("www.wzranked.com", "wzranked.com"))
	
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s", r.URL)
	})
	
	collector.OnHTML("a.items-center div", func(h *colly.HTMLElement) {
		WeaponStats.Guns = h.Text
	})

	// collector.OnHTML("a.text-custom-text-primary span", func(h *colly.HTMLElement) {
	// 	WeaponStats.SecHead = h.Text
	// })


	collector.OnHTML("th.text-custom-text-primary", func(h *colly.HTMLElement) {
		WeaponStats.SecHead = h.Text
	})

	collector.OnHTML("td.text-custom-text-secondary", func(h *colly.HTMLElement) {
		WeaponStats.WeaponTypePopularity = h.Text
	})


	// collector.OnHTML("parent class name text", func(h *colly.HTMLElement) {
	// 	selection := h.DOM

	// 	childNodes := selection.Children().Nodes
	// 	if len(childNodes) == 3 {
	// 		description := selection.Find("class name text").Text()
	// 		value := selection.FindNodes(childNodes[2]).Text()

	// 		fmt.Printf("%s: %s\n", description, value)
	// 	}


	// })


	collector.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	}) 

	collector.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		enc.Encode(WeaponStats)
	})

	collector.Visit(scrapeUrl)
}


// <a class="inline-flex items-center space-x-1 align-middle hover:underline hover:brightness-90" href="/wz2/meta/guns/cronen-squall"><div class="w-12 flex-shrink-0"><img alt="Cronen Squall" src="https://wzranked.b-cdn.net/images/mw2/guns/cronen-squall.png" width="48" height="24" decoding="async" data-nimg="1" loading="lazy" style="color: transparent;"></div><div>Cronen Squall</div></a>