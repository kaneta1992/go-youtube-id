package youtubeid

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

type Client struct {
	service *youtube.Service
	getMax  int64
}

func NewClient(key string, max int64) *Client {
	c := &Client{getMax: max}

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

func (c *Client) GetVideos(query string) Videos {
	videos := Videos{}
	call := c.service.Search.List("id,snippet").
		Q(query).
		MaxResults(c.getMax).
		Type("video")
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, NewVideo(item.Id.VideoId, item.Snippet.Title, nil, c))
		}
	}
	return videos
}

func (c *Client) GetRelatedVideos(v *Video) Videos {
	videos := Videos{}
	call := c.service.Search.List("id,snippet").
		RelatedToVideoId(v.id).
		MaxResults(c.getMax).
		Type("video")
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, NewVideo(item.Id.VideoId, item.Snippet.Title, v, c))
		}
	}
	return videos
}
