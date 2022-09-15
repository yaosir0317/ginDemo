package redisDB

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Article struct {
	ArticleId string `json:"articleId"`
	ArticleType string `json:"articleType"`
	AuthorId string `json:"authorId"`
	AuthorType string `json:"authorType"`
	CaptureTime int64 `json:"captureTime"`
	CaptureUpdateTime int64 `json:"captureUpdateTime"`
	Category int64 `json:"category"`
	CategoryCode string `json:"categoryCode"`
	CommentNum int64 `json:"commentNum"`
	CompressKey string `json:"compressKey"`
	Deleted int64 `json:"deleted"`
	DownNum int64 `json:"downNum"`
	Duration int64 `json:"duration"`
	Id int64 `json:"id"`
	MediaList []MediaInfo `json:"mediaList"`
	MediaRatio int64 `json:"mediaRatio"`
	OriginKey string `json:"originKey"`
	PlayNum int64 `json:"playNum"`
	Rate int64 `json:"rate"`
	SecondCategory int64 `json:"secondCategory"`
	ShareNum int64 `json:"shareNum"`
	ShowUpNum int64 `json:"showUpNum"`
	Source string `json:"source"`
	SourceCategory int64 `json:"sourceCategory"`
	SourceCommentNum int64 `json:"sourceCommentNum"`
	SourceDownNum int64 `json:"sourceDownNum"`
	SourceSecondCategory int64 `json:"sourceSecondCategory"`
	SourceShareNum int64 `json:"sourceShareNum"`
	SourceTimeCreated int64 `json:"sourceTimeCreated"`
	SourceUpNum int64 `json:"sourceUpNum"`
	TimeCreated int64 `json:"timeCreated"`
	TimeLastUpdated int64 `json:"timeLastUpdated"`
	Title string `json:"title"`
	UpNum int64 `json:"upNum"`
	Verify int64 `json:"verify"`
}

type MediaInfo struct {
	CoverUrl string `json:"coverUrl"`
	DownloadUrl string `json:"downloadUrl"`
	MediaHeight int64 `json:"mediaHeight"`
	MediaType string `json:"mediaType"`
	MediaUrl string `json:"mediaUrl"`
	MediaWidth int64 `json:"mediaWidth"`
	VideoDuration int64 `json:"videoDuration"`

}

func (i *impl) GetKey() (string, error) {
	str, err := i.client.Get("current:ces:version:ces2").Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return str, nil
}

func (i *impl) MGetKey() ([]string, error) {
	var ret []string
	dataTuple, err := i.client.MGet("current:ces:version:ces2", "current:ces:version:ces2.1").Result()
	if err != nil {
		fmt.Println(err)
		return ret, err
	}
	fmt.Println(dataTuple, reflect.TypeOf(dataTuple), reflect.TypeOf(dataTuple).Name(), reflect.TypeOf(dataTuple).Kind())
	ret = append(ret, dataTuple[0].(string))
	ret = append(ret, dataTuple[1].(string))
	return ret, err
}

func (i *impl) GetRemovePreQ(d string) (*[]string, error) {
	key := fmt.Sprintf("article:new:preQ:del:%s", d)
	data, err := i.client.SMembers(key).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil
}

func (i *impl) HGetInfo(articleId string) error {
	s, err := i.client.HGet("articles", articleId).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("s", s)
	var a Article
	err = json.Unmarshal([]byte(s), &a)
	fmt.Println("a", a)
	return nil
}
