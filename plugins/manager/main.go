package manager

import (
	"github.com/miracle0609/wxbot/engine/pkg/log"
	"github.com/miracle0609/wxbot/engine/pkg/sqlite"
)

var db sqlite.DB

func init() {
	// 初始化数据库
	if err := sqlite.Open("data/plugins/manager/manager.db", &db); err != nil {
		log.Fatalf("open sqlite db failed: %v", err)
	}

	// 注册定时任务插件
	registerCronjob()

	// 注册命令插件
	registerCommand()
}
