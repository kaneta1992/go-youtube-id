package youtubeid

import (
	"fmt"
	"math/rand"
)

type Video struct {
	client    *Client
	id        string
	title     string
	keyword   string
	prevVideo *Video
}

func NewVideo(id, title, keyword string, prevVideo *Video, c *Client) *Video {
	return &Video{
		client:    c,
		id:        id,
		title:     title,
		keyword:   keyword,
		prevVideo: prevVideo,
	}
}

func (v *Video) URL() string {
	return "https://www.youtube.com/watch?v=" + v.id
}

func (v *Video) Next() *Video {
	videos := v.client.GetRelatedVideos(v, 3)
	return videos.Random()
}

func (v *Video) Prev() *Video {
	return v.prevVideo
}

type Videos []*Video

func (v Videos) IsEmpty() bool {
	return len(v) <= 0
}

func (v Videos) First() *Video {
	if v.IsEmpty() {
		return nil
	}
	return v[0]
}

func (v Videos) Last() *Video {
	if v.IsEmpty() {
		return nil
	}
	index := len(v)
	return v[index-1]
}

func (v Videos) Random() *Video {
	if v.IsEmpty() {
		return nil
	}
	index := rand.Intn(len(v))

	fmt.Println(v[index])
	return v[index]
}
