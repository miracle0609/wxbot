package damoku

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"database/sql"
	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
	"github.com/miracle0609/wxbot/engine/pkg/log"
	"github.com/miracle0609/wxbot/engine/pkg/sqlite"
	"github.com/mattn/go-sqlite3"
)

//go:embed data

var db sqlite.DB
var err string
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
	
	if err := sql.Open("sqlite3", engine.GetDataFolder()+"/moku.db"); err != nil {
		log.Fatalf("open sqlite db failed: %v", err)
	}
	sqlStmt := "CREATE TABLE IF NOT EXISTS xinmoku (name TEXT, answer TEXT);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
    fmt.Sprintf("%q: %s\n", err, sqlStmt)
		return
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
	xinmokuFile, err := os.Open("data/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer xinmokuFile.Close()

	// insert system moku answer
	scanner := bufio.NewScanner(xinmokuFile)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) == 2 {
			sqlStmt := "INSERT INTO xinmoku (name, answer) VALUES (?, ?)"
			_, err = db.Exec(sqlStmt, fields[0], fields[1])
			//_, _, err = db.Orm.Table("xinmoku").FirstOrCreate(&Xinmoku{Name: fields[0], Answer: fields[1]}, "name = ? and answer = ?", fields[0], fields[1])
			if err != nil {
				fmt.Sprintf("%q: %s\n", err, fields[0])
			}
		} else {
			fmt.Sprintf("Invalid line: %s\n", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
