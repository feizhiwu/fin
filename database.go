package fin

import (
	"github.com/feizhiwu/gs/albedo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type db struct {
	connects map[string]*gorm.DB
}

func connectDB() *db {
	db := new(db)
	db.connects = make(map[string]*gorm.DB)
	dbConf := Config("db")
	Assert(dbConf != nil, "db config is not set")
	if len(dbConf.(map[interface{}]interface{})) > 1 {
		for k, v := range dbConf.(map[interface{}]interface{}) {
			db.connect(albedo.MakeString(k), v.(map[interface{}]interface{}))
		}
	} else {
		db.connect("db", dbConf.(map[interface{}]interface{}))
	}
	return db
}

func (db *db) connect(name string, options map[interface{}]interface{}) {
	ms := make(MS)
	for k, v := range options {
		ms[albedo.MakeString(k)] = albedo.MakeString(v)
	}
	if ms["charset"] == "" {
		ms["charset"] = "utf8"
	}
	connect, err := gorm.Open(ms["datatype"], ms["username"]+":"+ms["password"]+"@tcp("+ms["hostname"]+")/"+ms["database"]+"?charset="+ms["charset"])
	errInfo := ""
	if err != nil {
		errInfo = err.Error()
	}
	Assert(err == nil, errInfo)
	//全局禁用表复数
	connect.SingularTable(true)
	db.connects[name] = connect
}
