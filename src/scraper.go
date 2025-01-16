package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Hello world")
	c := colly.NewCollector()
	const latitude = "43.35689"
	const longitude = "-3.0"
	c.Visit("https://es.wallapop.com/app/search?filters_source=category_navigation&latitude=" + latitude + "&" + "longitude=" + longitude)
}
