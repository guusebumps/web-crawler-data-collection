package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Title found:", e.Text)
	})
	c.Visit("https://github.com")
}
