package youtube

type Video struct {
	id        string
	prevVideo *Video
	nextVideo *Video
}

func NewVideo(id string, prevVideo *Video) *Video {
	return &Video{
		id:        id,
		prevVideo: prevVideo,
	}
}

func (v *Video) Next() *Video {
	if v.nextVideo == nil {
		// TODO: 次のビデオをここでスクレイピングする
		// 見つからなければnilを返す
		v.nextVideo = NewVideo(v.id, v)
	}
	return v.nextVideo
}

func (v *Video) Prev() *Video {
	if v.prevVideo == nil {
		return nil
	}
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
	// TODO: ランダムにする
	if v.IsEmpty() {
		return nil
	}
	index := 0
	return v[index]
}
