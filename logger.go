package fin

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strings"
	"time"
)

//*************************DB log start*************************

type dbLog struct {
	name string
	ip   string
	db   *gorm.DB
}

func newDBLog(name, ip string, db *gorm.DB) {
	db.LogMode(true)
	db.SetLogger(&dbLog{
		name,
		ip,
		db,
	})
}

func (l *dbLog) Print(values ...interface{}) {
	logFile := logFile()
	level := values[0]
	if level == "sql" {
		var (
			//source    = values[1]
			queryTime = values[2]
			sql       = values[3]
			param     = values[4]
		)
		if queryTime.(time.Duration) > time.Millisecond*300 {
			level = "[SLOW]"
		} else {
			level = "[OK]"
		}
		log.New(logFile, "", log.LstdFlags).Printf("%s[%s]%s %s %s in %s", l.ip, l.name, level, sql, param, queryTime)
		if Mode() == DebugMode {
			log.Printf("%s[%s]%s %s %s in %s", l.ip, l.name, level, sql, param, queryTime)
		}
	} else {
		log.New(logFile, "", log.LstdFlags).Printf("%s[%s]%s", l.ip, l.name, values)
		if Mode() == DebugMode {
			log.Printf("%s[%s]%s", l.ip, l.name, values)
		}
	}
}

//*************************DB log end*************************

//日志middleware
func logger(c *Context) {
	t := time.Now()
	logFile := logFile()
	for _, v := range dbs {
		newDBLog(strings.ToUpper(v), c.ClientIP(), BeginDB(v))
	}
	c.Next()
	for _, v := range dbs {
		CommitDB(v)
	}
	log.New(logFile, "", log.LstdFlags).Printf("%s[%d] %s in %v\n", c.ClientIP(), c.StatusCode, c.Request.RequestURI, time.Since(t))
	log.New(logFile, "", log.LstdFlags).Printf("%s", "----------------------------------------------------------------------")
	if Mode() == DebugMode {
		log.Printf("%s[%d] %s in %v\n", c.ClientIP(), c.StatusCode, c.Request.RequestURI, time.Since(t))
		log.Printf("%s", "----------------------------------------------------------------------")
	}
}

func logFile() *os.File {
	//创建文件路径
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, os.ModePerm)
	}
	var logFix string
	if Mode() == TestMode {
		logFix = "-test.log"
	} else {
		logFix = ".log"
	}
	filePath := logPath + "/" + time.Now().Format("02") + logFix
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
	}
	logFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	return logFile
}
