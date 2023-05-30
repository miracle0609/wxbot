package chatgpt

import (
	"path/filepath"
	"time"

	"github.com/miracle0609/wxbot/engine/pkg/log"
	"github.com/miracle0609/wxbot/engine/pkg/utils"
	"github.com/miracle0609/wxbot/engine/robot"
)

// 设置图片相关指令
func setImageCommand(ctx *robot.Ctx, msg string, command string) {
	switch command {
	case "作画":
		b64, err := AskChatGptWithImage(ctx, msg, time.Second)
		if err != nil {
			log.Errorf("出错了稍后再试，Err：%s", err.Error())
			ctx.ReplyTextAndAt("提问频率过快，请稍后再试")
			return
		}
		filename := filepath.Join("data/plugins/chatgpt/cache", msg+".png")
		if err := utils.Base64ToImage(b64, filename); err != nil {
			log.Errorf("作画失败，Err: %s", err.Error())
			ctx.ReplyTextAndAt("作画失败，请重试")
			return
		}
		ctx.ReplyImage("local://" + filename)
	}
}
