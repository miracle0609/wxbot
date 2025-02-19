package diange

import (
	"fmt"
	"net/url"
	"strconv"
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
)

type AutoGenerated struct {
	Result Result `json:"result"`
	Code   int    `json:"code"`
}
type Artists struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	PicURL    interface{}   `json:"picUrl"`
	Alias     []interface{} `json:"alias"`
	AlbumSize int           `json:"albumSize"`
	PicID     int           `json:"picId"`
	FansGroup interface{}   `json:"fansGroup"`
	Img1V1URL string        `json:"img1v1Url"`
	Img1V1    int           `json:"img1v1"`
	Trans     interface{}   `json:"trans"`
}
type Artist struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	PicURL    interface{}   `json:"picUrl"`
	Alias     []interface{} `json:"alias"`
	AlbumSize int           `json:"albumSize"`
	PicID     int           `json:"picId"`
	FansGroup interface{}   `json:"fansGroup"`
	Img1V1URL string        `json:"img1v1Url"`
	Img1V1    int           `json:"img1v1"`
	Trans     interface{}   `json:"trans"`
}
type Album struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Artist      Artist `json:"artist"`
	PublishTime int64  `json:"publishTime"`
	Size        int    `json:"size"`
	CopyrightID int    `json:"copyrightId"`
	Status      int    `json:"status"`
	PicID       int64  `json:"picId"`
	Mark        int    `json:"mark"`
}
type Songs struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Artists     []Artists     `json:"artists"`
	Album       Album         `json:"album"`
	Duration    int           `json:"duration"`
	CopyrightID int           `json:"copyrightId"`
	Status      int           `json:"status"`
	Alias       []interface{} `json:"alias"`
	Rtype       int           `json:"rtype"`
	Ftype       int           `json:"ftype"`
	Mvid        int           `json:"mvid"`
	Fee         int           `json:"fee"`
	RURL        interface{}   `json:"rUrl"`
	Mark        int           `json:"mark"`
}
type Result struct {
	Songs     []Songs `json:"songs"`
	HasMore   bool    `json:"hasMore"`
	SongCount int     `json:"songCount"`
}


func init() {
	engine := control.Register("diange", &control.Options{
		Alias: "点歌",
		Help: "指令:\n" +
			"* 点歌 [歌名]\n" ,
	})

	engine.OnRegex(`^点歌 ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		word := ctx.State["regex_matched"].([]string)[1]
		fmt.Printf("歌曲 = %s\n", word)
		if testx, err := getSong(word); err == nil {
			if testx == nil {
				ctx.ReplyText("出错了，稍后尝试")
			} else {
			 for _, song := range testx.Result.Songs {
				songurl := song.ID
				fmt.Printf("歌曲id = %d\n", songurl)
				geurl := "https://music.163.com/song/media/outer/url?id="
				geurl += strconv.Itoa(songurl)
				geurl += ".mp3"
				fmt.Printf("歌曲连接 = %s\n", geurl)
				for _, artist := range song.Artists {
					ctx.ReplyMusic(song.Name, artist.Name, "网易云/wx8dd6ecd81906fd84", "http://music.163.com/song/media/outer/", geurl, artist.Img1V1URL)
					break
				}
				break
			}
				//ReplyMusic(name, author, app, jumpUrl, musicUrl, coverUrl string)
			}
		}
	})
}


func getSong(keyword string)(*AutoGenerated, error) {
	var resp AutoGenerated
	api := "http://64.112.43.106:3000/search?keywords=" + url.QueryEscape(keyword)
	if err := req.C().SetBaseURL(api).Post().Do().Into(&resp); err != nil {
			fmt.Println(resp)
			return nil, err
	}
	if resp.Code != 200 {
		return nil, nil
	}
	fmt.Printf("之后\n")
	fmt.Println(resp)
	
	return &resp, nil
}
