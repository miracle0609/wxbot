package youdaofanyi

import (
	"fmt"
	"time"
	"net/url"
	"strconv"
	"github.com/imroc/req/v3"

	"github.com/yqchilde/wxbot/engine/control"
	"github.com/yqchilde/wxbot/engine/robot"
)

func init() {
	engine := control.Register("caimiyu", &control.Options{
		Alias: "猜谜语",
		Help: "指令:\n" +
			"* 猜谜语\n" ,
	})

	engine.OnRegex(`(^猜谜语) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		word := ctx.State["regex_matched"].([]string)
		userIssue := ctx.MessageString()
		recv, cancel := ctx.EventChannel(ctx.CheckUserSession()).Repeat()
		defer cancel()
		if data, err := getZiMi(word); err == nil {
			if data == nil {
				ctx.ReplyText("出错了，请稍后尝试")
			} else {
				ctx.ReplyText(fmt.Sprintf("🔎 题目:(60秒之后自动给出答案）\n %s", (data.Result.riddle).String()+","(data.Result.type).String())
				timeLimit := time.After(60 * time.Second)
				for {
					select {
					case <-timeLimit:
						ctx.ReplyTextAndAt(fmt.Sprintf("🔎 时间到,正确答案是：\n %s", (data.Result.answer).String()))
						return
					case ctx := <-recv:
						userAnswer := ctx.MessageString()
						if userAnswer == data.Result.answer {
							ctx.ReplyText("恭喜你，回答正确,猜谜结束")
							return
						}
						ctx.ReplyTextAndAt(fmt.Sprintf("很遗憾，你回答错误"))
						return
					}
				}
			}
		} else {
			ctx.ReplyText("查询失败，这一定不是bug🤔")
		}
	})
}

type apiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result string `json:"result"`
}

func getZiMi(keyword string) (*apiResponse, error) {
	var data apiResponse
	api := "https://api.qqsuu.cn/api/dm-caizimi"
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		return nil, err
	}
	if len(data.Result) == 0 {
		return nil, nil
	}
	return &data, nil
}
