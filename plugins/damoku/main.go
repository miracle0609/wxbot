package damoku

import (
	"embed"
	"errors"
	"fmt"
	"strings"
	"bufio"
	"github.com/imroc/req/v3"

	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
	"github.com/miracle0609/wxbot/engine/pkg/log"
	"github.com/miracle0609/wxbot/engine/pkg/sqlite"
)

//go:embed data
var mokuData embed.FS

type Xinmoku struct {
	Name    string    `gorm:"column:name;index"`    // 
	Answer    string `gorm:"column:answer;index"`    // 
}

func init() {
	engine := control.Register("damoku", &control.Options{
		Alias: "打魔窟",
		Help: "指令:\n" +
			"* 打[毒瘤名字/毒瘤组合编号]\n" ,
	})
	
	if err := sqlite.Open(engine.GetDataFolder()+"/moku.db", &db); err != nil {
		log.Fatalf("open sqlite db failed: %v", err)
	}
	if err := db.Create("xinmoku", &Xinmoku{}); err != nil {
		log.Fatalf("create xinmoku table failed: %v", err)
	}
	// 导入新魔窟打法
	initXinmoku()

	engine.OnRegex(`(^打) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		word := ctx.State["regex_matched"].([]string)[1]
		fmt.Printf("毒瘤名字/毒瘤组合编号 = %s\n", word)
		var answer string
		db.Orm.Table("xinmoku").Select("answer").Where("name = ?", word).Scan(&answer)
		ctx.ReplyTextAndAt(fmt.Sprintf("🔎 %s", answer))
	})
}


func initXinmoku() {
	xinmokuFile, err := mokuData.ReadFile("data/xinmoku.txt")
	if err != nil {
		log.Errorf("[Xinmoku] 获取魔窟答案失败, error:%s", err.Error())
		return
	}
	defer xinmokuFile.Close()
	// insert system moku answer
	scanner := bufio.NewScanner(xinmokuFile)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) == 2 {
			//sqlStmt := "INSERT INTO xinmoku (name, answer) VALUES (?, ?)"
			//_, err = db.Exec(sqlStmt, fields[0], fields[1])
			_, err = db.Orm.Table("xinmoku").FirstOrCreate(&Xinmoku{Name: fields[0], Answer: fields[1]}, "name = ? and answer = ?", fields[0], fields[1])
			if err != nil {
				fmt.Sprintf("%q: %s\n", err, sqlStmt)
			}
		} else {
			fmt.Sprintf("Invalid line: %s\n", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
