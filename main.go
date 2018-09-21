package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.youtube.com/results?search_query=%E5%8A%A0%E8%97%A4%E7%B4%94%E4%B8%80&sp=EgIQAQ%253D%253D")
	if err != nil {
		panic(err)
	}

	selection := doc.Find("h3 a")
	cnt := 0
	selection.Each(func(index int, s *goquery.Selection) {
		fmt.Println(s.AttrOr("href", "not exists"))
		cnt++
	})
	fmt.Println(cnt)
}
