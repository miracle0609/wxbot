package vlw

import (
	"fmt"
	"os"
	"time"

	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/pkg/log"
)

type MessageResp struct {
	Code      int    `json:"Code"`
	Result    string `json:"Result"`
	ReturnStr string `json:"ReturnStr"`
	ReturnInt string `json:"ReturnInt"`
}

func NewRequest() *req.Client {
	c := req.C().
		SetLogger(log.GetLogger()).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			if os.Getenv("DEBUG") == "true" {
				client.DevMode()
			}
			return nil
		}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				if dump := resp.Dump(); dump != "" {
					resp.Err = fmt.Errorf("%s\nraw content:\n%s", resp.Err.Error(), resp.Dump())
				}
				return nil
			}
			var dataResp MessageResp
			if err := resp.Into(&dataResp); err != nil {
				resp.Err = fmt.Errorf("解析Response失败, error: %s", err.Error())
				return nil
			}
			if dataResp.Code != 0 {
				resp.Err = fmt.Errorf(resp.String())
				return nil
			}
			return nil
		})
	return c
}
