package youtubeid

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

type Client struct {
	service *youtube.Service
}

func NewClient(key string) *Client {
	c := &Client{}

	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	var err error
	c.service, err = youtube.New(client)
	if err != nil {
		panic(err)
	}

	return c
}

func (c *Client) GetVideos(query string, max int64) Videos {
	videos := Videos{}
	call := c.service.Search.List("id,snippet").
		Q(query).
		MaxResults(max).
		Type("video")
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, NewVideo(item.Id.VideoId, item.Snippet.Title, query, nil, c))
		}
	}
	return videos
}

func (c *Client) GetRelatedVideos(v *Video, max int64) Videos {
	videos := Videos{}
	call := c.service.Search.List("snippet").
		RegionCode("JP").
		RelatedToVideoId(v.id).
		MaxResults(max).
		Type("video")
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, NewVideo(item.Id.VideoId, item.Snippet.Title, v.keyword, v, c))
		}
	}
	return videos
}
