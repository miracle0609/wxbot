package diange

import (
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
)

type Song struct {
    Name       string `json:"name"`
    ID         int    `json:"id"`
    Artists    []struct {
        ID      int      `json:"id"`
        Name    string   `json:"name"`
        Tns     []string `json:"tns"`
        Alias   []string `json:"alias"`
        Alia    []string `json:"alia"`
    } `json:"ar"`
    Album      struct {
        ID       int    `json:"id"`
        Name     string `json:"name"`
        PicUrl   string `json:"picUrl"`
        Tns      []string `json:"tns"`
        PicStr   string `json:"pic_str"`
        Pic      int    `json:"pic"`
    } `json:"al"`
    Duration   int    `json:"dt"`
    High       struct {
        Br      int    `json:"br"`
        Fid     int    `json:"fid"`
        Size    int    `json:"size"`
        Vd      int    `json:"vd"`
        Sr      int    `json:"sr"`
    } `json:"h"`
    Middle     struct {
        Br     int    `json:"br"`
        Fid     int    `json:"fid"`
        Size    int    `json:"size"`
        Vd      int    `json:"vd"`
        Sr      int    `json:"sr"`
    } `json:"m"`
    Low        struct {
        Br      int    `json:"br"`
        Fid     int    `json:"fid"`
        Size    int    `json:"size"`
        Vd      int    `json:"vd"`
        Sr      int    `json:"sr"`
    } `json:"l"`
    SuperQuality struct {
        Br      int    `json:"br"`
        Fid     int    `json:"fid"`
        Size    int    `json:"size"`
        Vd      int    `json:"vd"`
        Sr      int    `json:"sr"`
    } `json:"sq"`
    Privilege  struct {
        ID        int `json:"id"`
        Fee       int `json:"fee"`
        Payed     int `json:"payed"`
        Pl        int `json:"pl"`
        Dl        int `json:"dl"`
        Sp        int `json:"sp"`
        Cp        int `json:"cp"`
        Subp      int `json:"subp"`
        Cs        bool `json:"cs"`
        Maxbr     int `json:"maxbr"`
        Fl        int `json:"fl"`
        Toast     bool`json:"toast"`
        Flag      int `json:"flag"`
        PreSell   bool `json:"preSell"`
        PlayMaxbr int `json:"playMaxbr"`
        DownloadMaxbr int `json:"downloadMaxbr"`
        MaxBrLevel string `json:"maxBrLevel"`
        PlayMaxBrLevel string `json:"playMaxBrLevel"`
        DownloadMaxBrLevel string `json:"downloadMaxBrLevel"`
        PlLevel   string `json:"plLevel"`
        DlLevel   string `json:"dlLevel"`
        FlLevel   string `json:"flLevel"`
        Rscl      interface{} `json:"rscl"`
        FreeTrialPrivilege struct {
            ResConsumable bool `json:"resConsumable"`
            UserConsumable bool `json:"userConsumable"`
            ListenType interface{} `json:"listenType"`
        } `json:"freeTrialPrivilege"`
        ChargeInfoList []struct {
            Rate int `json:"rate"`
            ChargeUrl interface{} `json:"chargeUrl"`
            ChargeMessage interface{} `json:"chargeMessage"`
            ChargeType int `json:"chargeType"`
        } `json:"chargeInfoList"`
    } `json:"privilege"`
}

type Result struct {
    SearchQcReminder interface{} `json:"searchQcReminder"`
    Songs             []Song `json:"songs"`
    SongCount         int    `json:"songCount"`
}


func init() {
	engine := control.Register("diange", &control.Options{
		Alias: "点歌",
		Help: "指令:\n" +
			"* 点歌 [歌名]\n" ,
	})

	engine.OnRegex(`^点歌 ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		word := ctx.State["regex_matched"].([]string)[1]
		if testx, err := getSong(word); err == nil {
			if testx == nil {
				ctx.ReplyText("出错了，稍后尝试")
			} else {
				songurl := testx.Songs[0].ID
				geurl := "https://music.163.com/song/media/outer/url?id="
				geurl += songurl
				gerul += ".mp3"
				ctx.ReplyMusic(testx.Songs[0].Name, testx.Songs[0].Artists[0].Name, "网易云/wx8dd6ecd81906fd84", "http://music.163.com/song/media/outer/", geurl , testx.Songs[0].Album.PicUrl)
				//ReplyMusic(name, author, app, jumpUrl, musicUrl, coverUrl string)
			}
		}
	})
}


func getSong(keyword string)(*Result, error) {
	var resp Result
	api := "http://64.112.43.106:3000/search?keywords=" + keyword + "&limit=1"
	if err := req.C().SetBaseURL(api).Get().Do().Into(&resp); err != nil {
			return nil, err
	}
	return &resp, nil
}
