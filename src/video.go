package youtubeid

type Video struct {
	client    *Client
	id        string
	title     string
	prevVideo *Video
}

func NewVideo(id, title string, prevVideo *Video, c *Client) *Video {
	return &Video{
		client:    c,
		id:        id,
		title:     title,
		prevVideo: prevVideo,
	}
}

func (v *Video) URL() string {
	return "https://www.youtube.com/watch?v=" + v.id
}

func (v *Video) Next() *Video {
	// TODO: 次のビデオをここでスクレイピングする
	// 見つからなければnilを返す
	return NewVideo(v.id, v.title, v, v.client)
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
	// TODO: ランダムにする
	if v.IsEmpty() {
		return nil
	}
	index := 0
	return v[index]
}
