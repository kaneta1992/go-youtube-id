package main

import (
	"fmt"
	"os"

	youtube "github.com/kaneta1992/go-youtube-id/src"
)

func main() {
	c := youtube.NewClient(os.Getenv("YOUTUBE_KEY"), 50)
	videos := c.GetVideos("加藤純一")
	fmt.Println(videos.First())
	fmt.Println(videos.Last())
	for _, v := range videos {
		fmt.Println(v.URL())
	}

	videos = c.GetRelatedVideos(videos.First())
	fmt.Println(videos.First())
	fmt.Println(videos.Last())
	for _, v := range videos {
		fmt.Println(v.URL())
	}
}
