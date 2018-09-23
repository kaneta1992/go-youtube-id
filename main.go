package main

import (
	"fmt"

	youtube "github.com/kaneta1992/go-youtube-id/src"
)

func main() {
	c := youtube.NewClient("加藤純一")
	fmt.Println(c.GetVideos().First())
	fmt.Println(c.GetVideos().Last())
	c.GetVideos().First().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next()

	for _, v := range c.GetVideos() {
		fmt.Println(v.URL())
	}
}
