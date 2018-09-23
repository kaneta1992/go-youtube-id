package youtube

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Video struct {
	id           string
	title        string
	prevVideo    *Video
	relatedCache Videos
}

func NewVideo(id, title string, prevVideo *Video) *Video {
	return &Video{
		id:        id,
		title:     title,
		prevVideo: prevVideo,
	}
}

func (v *Video) URL() string {
	return "https://www.youtube.com/watch?v=" + v.id
}

func (v *Video) updateRelatedCache() {
	time.Sleep(time.Microsecond * 100)
	if v.relatedCache == nil {
		doc, err := goquery.NewDocument(v.URL())
		if err != nil {
			panic(err)
		}
		/*
			<div class="content-wrapper">
				<a href="/watch?v=GkfN-M3xtu8" class=" content-link spf-link  yt-uix-sessionlink      spf-link " data-sessionlink="itct=CCkQpDAYACITCIWjrajk0N0CFUTJWAodR4YA5yj4HTIHYXV0b25hdkjB2Ojr_uKmpNQB"  title="17時間半PUBG配信の最期【2018/01/05】" rel=" spf-prefetch nofollow" data-visibility-tracking="CCkQpDAYACITCIWjrajk0N0CFUTJWAodR4YA5yj4HUDv7cbvjL_zoxo=" >
				  <span dir="ltr" class="title" aria-describedby="description-id-259358">
				    タイトル
				  </span>
				  <span class="accessible-description" id="description-id-259358">
				     - 長さ: 23:00。
				  </span>
				  <span class="stat attribution"><span class="" >チャンネル名</span></span>
				  <span class="stat view-count">視聴回数 26,693 回</span>
				</a>
			</div>
		*/
		selection := doc.Find("div.content-wrapper a")
		selection.Each(func(index int, s *goquery.Selection) {
			if id, ok := s.Attr("href"); ok {
				// "/watch?v="を切り取る
				id = id[9:]
				title := s.Find("span").First().Text()
				title = strings.TrimSpace(title)
				v.relatedCache = append(v.relatedCache, NewVideo(id, title, v))
			}
		})
	}
}

func (v *Video) Next() *Video {
	v.updateRelatedCache()
	// TODO: 次の動画をすでに再生したことがある場合は関連動画にする
	fmt.Println(v.relatedCache.First())
	return v.relatedCache.First()
}

func (v *Video) Relate() *Video {
	v.updateRelatedCache()
	video := v.relatedCache.Random()
	fmt.Println(video)
	return video
}

func (v *Video) Prev() *Video {
	return v.prevVideo
}

// TODO: 関連動画とかもあるといいかも

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

	return v[index]
}
