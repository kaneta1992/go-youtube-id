package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	youtube "github.com/kaneta1992/go-youtube-id/src"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := youtube.NewClient(os.Getenv("YOUTUBE_KEY"))
	videos := c.GetVideos("加藤純一", 10)
	fmt.Println(videos.First())
	fmt.Println(videos.Last())
	for _, v := range videos {
		fmt.Println(v.URL())
	}

	next := videos.Random().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next()

	fmt.Println(next.URL())
}
