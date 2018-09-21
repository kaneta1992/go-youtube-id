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
