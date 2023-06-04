package damoku

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"database/sql"
	"regexp"
	"strconv"
	"github.com/miracle0609/wxbot/engine/control"
	"github.com/miracle0609/wxbot/engine/robot"
	"github.com/miracle0609/wxbot/engine/pkg/log"
	_ "github.com/mattn/go-sqlite3"
)

type Xinmoku struct {
	Name    int   `db:"name"`
	Answer    string `db:"answer"`
}



func init() {
	engine := control.Register("damoku", &control.Options{
		Alias: "打魔窟",
		Help: "指令:\n" +
			"* 打[毒瘤组合编号],毒瘤[毒瘤名字]\n" ,
	})
	//initDatabase()
	
	
	engine.OnRegex(`(^打组合) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		db, err := sql.Open("sqlite3", "E:/robotTestProject/wxbot/data/plugins/damoku/moku.db")
		if err != nil {
			log.Fatal(err)
		}
		word := ctx.State["regex_matched"].([]string)[2]
		fmt.Printf("毒瘤名字/毒瘤组合编号 = %s\n", word)
		re := regexp.MustCompile(`\d+`)
		word = re.FindString(word)
		num, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		fmt.Printf("字符串 %s 转换为整数后为 %d\n", word, num)
		sqlStm := "SELECT * FROM xinmoku where name = ?"
		result, err := db.Query(sqlStm, word)
		if err != nil {
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
			result.Close()
			db.Close()
			fmt.Sprintf("error: %s\n", err)	
			return
		}
		var found bool
		for result.Next() {
			var name string
			var answer string
			if err := result.Scan(&name, &answer); err != nil {
				fmt.Sprintf("error: %s\n", err)
				ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
				result.Close()
				db.Close()
				return
			}
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 %s", answer))
			result.Close()
		    found = true
			db.Close()
			return
	    }
		if !found {
			result.Close()
			db.Close()
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
			return
		}
		
	})
	engine.OnRegex(`(^毒瘤) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		db, err := sql.Open("sqlite3", "E:/robotTestProject/wxbot/data/plugins/damoku/moku.db")
		if err != nil {
			log.Fatal(err)
		}
		word := ctx.State["regex_matched"].([]string)[2]
		fmt.Printf("毒瘤名字 = %s\n", word)
		myMap := map[string]int{"克里斯蒂娜": 1000,"暗刺": 1001,"蒂娜": 1002,"塔洛斯将军": 1003,"将军": 1004,"海拉伯爵": 1005,"海拉": 1006,"泰坦T9": 1007,"泰坦": 1008,"钢蛋": 1009,"胖子": 1010,"艾莉瑞尔": 1011,"骚瑞": 1012,"玛格丽特": 1013,"玛丽": 1014,"玛格": 1015,"光法": 1016,"卡斯兰": 1017,"暗法": 1018,"蕾贝卡": 1019,"绿战": 1020,"哈迪斯": 1021,"毛驴": 1022,"骨牧": 1023,"王": 1024,"维塔斯博士": 1025,"博士": 1026,"烈": 1027,"阿努比斯": 1028,"狗头": 1029,"莱戈拉王子": 1030,"莱戈拉": 1031,"王子": 1032,"施莱德": 1033,"红刺": 1034,"艾芙琳": 1035,"光奶": 1036,"诺娃": 1037,"弱娃": 1038,"梅": 1039,"莎拉芙": 1040,"沙拉芙": 1041,"光弓": 1042,"光攻": 1043,"特蕾莎": 1044,"特傻": 1045,"岚": 1046,"图雅": 1047,"狄安娜": 1048,"希尔芙": 1049,"巴鲁贝尔": 1050,"狗熊": 1051,"奥黛丽夫人": 1052,"夫人": 1053,"瑰月": 1054,"鬼月": 1055,"奥利维亚": 1056,"红牧": 1057,"泷": 1058,"格莫瑞": 1059,"骨刺": 1060,"丝西娜": 1061,"克尔伯鲁": 1062,"棺材": 1063,"茉斯拉": 1064,"玛瑟伊尔": 1065,"甲虫": 1066,"塔萨王子": 1067,"布雷泽": 1068,"黑炮": 1069,"塔纳托斯": 1070,"稻草人": 1071,"兰斯特": 1072,"鸟人": 1073}
		num := myMap[word]
		fmt.Printf("数据库编号 = %d\n", num)
		sqlStm := "SELECT * FROM xinmoku where name = ?"
		result, err := db.Query(sqlStm, num)
		if err != nil {
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
			result.Close()
			db.Close()
			fmt.Sprintf("error: %s\n", err)	
			return
		}
		var found bool
		for result.Next() {
			var name string
			var answer string
			if err := result.Scan(&name, &answer); err != nil {
				fmt.Sprintf("error: %s\n", err)
				ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
				result.Close()
				db.Close()
				return
			}
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 %s", answer))
			result.Close()
		    found = true
			db.Close()
			return
	    }
		if !found {
			result.Close()
			db.Close()
			ctx.ReplyTextAndAt(fmt.Sprintf("🔎 暂未查到"))
			return
		}
		
	})
}

func initDatabase() {
	db, err := sql.Open("sqlite3", "E:/robotTestProject/wxbot/data/plugins/damoku/moku.db")
    if err != nil {
        log.Fatal(err)
    }
	sqlStmt := "CREATE TABLE IF NOT EXISTS xinmoku (name int, answer TEXT);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Sprintf("%q: %s\n", err, sqlStmt)
		return
	}
	// 导入新魔窟打法
	xinmokuFile, err := os.Open("E:\\robotTestProject\\wxbot\\plugins\\damoku\\xinmoku.txt")
	if err != nil {
		log.Fatal(err)
	}
	// insert system moku answer
	scanner := bufio.NewScanner(xinmokuFile)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) == 2 {
			sqlStmt := "INSERT INTO xinmoku (name, answer) VALUES (?, ?)"
			_, err = db.Exec(sqlStmt, fields[0], fields[1])
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
	xinmokuFile.Close()
	db.Close()
}
