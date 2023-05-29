package wangyiyunRandSong

import (
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
)

type ApiResponse struct {
	Code   int    `json:"code"`
	Data struct {
		Name       string `json:"name"`
		Url       string `json:"url"`
		Picurl      string `json:"picurl"`
		Artistsname  string `json:"artistsname"`
	} `json:"data"`
}

func init() {
	engine := control.Register("caimiyu", &control.Options{
		Alias: "网易云",
		Help: "指令:\n" +
			"* 网易云\n" ,
	})

	engine.OnRegex(`(^网易云) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		if testx, err := getSong(); err == nil {
			if testx == nil {
				ctx.ReplyText("出错了，稍后尝试")
			} else {
				ctx.ReplyMusic(testx.Data.Name, testx.Data.Artistsname, "网易云/wx8dd6ecd81906fd84", "http://music.163.com/song/media/outer/", testx.Data.Url, testx.Data.Picurl)
				//ReplyMusic(name, author, app, jumpUrl, musicUrl, coverUrl string)
			}
		}
	})
}


func getSong()(*ApiResponse, error) {
	var resp ApiResponse
	api := "https://api.uomg.com/api/rand.music?sort=热歌榜&format=json"
	if err := req.C().SetBaseURL(api).Get().Do().Into(&resp); err != nil {
			return nil, err
	}
	if resp.Code != 1 {
		return nil, nil
	}
	return &resp, nil
}
