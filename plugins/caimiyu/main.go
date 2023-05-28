package caimiyu

import (
	"fmt"
	"time"
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
)

type apiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Riddle       string `json:"riddle"`
		Answer       string `json:"answer"`
		Disturb      string `json:"disturb"`
		Description  string `json:"description"`
		Type         string `json:"type"`
	} `json:"result"`
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
		if data, err := getZiMi(); err == nil {
			if data == nil {
				ctx.ReplyText("出错了，请稍后尝试")
			} else {
				ctx.ReplyText(fmt.Sprintf("🔎 题目:60秒之后自动给出答案\n %s", data.Result.Riddle+","+data.Result.Type))
				timeLimit := time.After(60 * time.Second)
				for {
					select {
					case <-timeLimit:
						ctx.ReplyTextAndAt(fmt.Sprintf("🔎 时间到,正确答案是：\n %s", data.Result.Answer))
						return
					case ctx := <-recv:
						userAnswer := ctx.MessageString()
						if userAnswer == data.Result.Answer {
							ctx.ReplyText("恭喜你，回答正确,猜谜结束")
							return
						}
						ctx.ReplyTextAndAt("很遗憾，你回答错误")
						return
					}
				}
			}
		} else {
			ctx.ReplyText("查询失败，这一定不是bug🤔")
		}
	})
}


func getZiMi() (*apiResponse, error) {
	var data apiResponse
	api := "https://api.qqsuu.cn/api/dm-caizimi"
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
