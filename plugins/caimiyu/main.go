package caimiyu

import (
	"fmt"
	"time"
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
)

type ApiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data struct {
		Riddle       string `json:"riddle"`
		Answer       string `json:"answer"`
		Disturb      string `json:"disturb"`
		Description  string `json:"description"`
		Type         string `json:"type"`
	} `json:"data"`
}

func init() {
	engine := control.Register("caimiyu", &control.Options{
		Alias: "猜谜语",
		Help: "指令:\n" +
			"* 猜谜语\n" ,
	})

	engine.OnRegex(`(^猜谜语) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		recv, cancel := ctx.EventChannel(ctx.CheckUserSession()).Repeat()
		defer cancel()
		if testx, err := getZiMi(); err == nil {
			if testx == nil {
				ctx.ReplyText("我还不会，稍后尝试")
			} else {
				ctx.ReplyText(fmt.Sprintf("🔎 题目:60秒之后自动给出答案\n %s", testx.Data.Riddle+","+testx.Data.Type))
				timeLimit := time.After(60 * time.Second)
				for {
					select {
					case <-timeLimit:
						ctx.ReplyTextAndAt(fmt.Sprintf("🔎 时间到,正确答案是：\n %s", testx.Data.Answer))
						return
					case ctx := <-recv:
						userAnswer := ctx.MessageString()
						if userAnswer == testx.Data.Answer {
							ctx.ReplyText("恭喜你，回答正确,猜谜结束")
							return
						} else if userAnswer != testx.Data.Answer {
							ctx.ReplyTextAndAt("很遗憾，你回答错误")
						}
					}
				}
			}
	})
}


func getZiMi()(*apiResponse, error) {
	var resp ApiResponse
	api := "https://api.qqsuu.cn/api/dm-caizimi"
	if err := req.C().SetBaseURL(api).Get().Do().Into(&resp); err != nil {
			return nil, err
	}
	if resp.Code != 200 {
		return nil, err
	}
	return &resp, nil
}
