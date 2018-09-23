package youtube

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	videos Videos
}

func NewClient(query string) *Client {
	c := &Client{}
	encoded := url.QueryEscape(query)
	doc, err := goquery.NewDocument("https://www.youtube.com/results?search_query=" + encoded + "&sp=EgIQAQ%253D%253D")
	if err != nil {
		panic(err)
	}

	selection := doc.Find("h3 a")
	selection.Each(func(index int, s *goquery.Selection) {
		if id, ok := s.Attr("href"); ok {
			// "/watch?v="を切り取る
			id = id[9:]
			title := s.Text()
			c.videos = append(c.videos, NewVideo(id, title, nil))
		}
	})
	return c
}

func (c *Client) GetVideos() Videos {
	return c.videos
}
