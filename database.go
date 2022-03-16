package fin

import (
	"github.com/feizhiwu/gs/albedo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbc MG            //db connects
var cdb map[uint64]MG //context db
var dbs []string

func DBConnects() MG {
	return dbc
}

// BeginDB 开启一个事务
//此为妥协做法，官方不建议获取协程id
func BeginDB(name string) *gorm.DB {
	if _, ok := cdb[GetGoroutineID()]; !ok {
		cdb[GetGoroutineID()] = make(MG)
	}
	if _, ok := cdb[GetGoroutineID()][name]; !ok {
		cdb[GetGoroutineID()][name] = dbc[name].Begin()
	}
	return cdb[GetGoroutineID()][name]
}

func CommitDB(name string) {
	if cdb[GetGoroutineID()][name] != nil {
		cdb[GetGoroutineID()][name].Commit()
		delete(cdb[GetGoroutineID()], name)
	}
	if len(cdb[GetGoroutineID()]) == 0 {
		delete(cdb, GetGoroutineID())
	}
}

func RollbackDB(name string) {
	if cdb[GetGoroutineID()][name] != nil {
		cdb[GetGoroutineID()][name].Rollback()
		delete(cdb[GetGoroutineID()], name)
	}
	if len(cdb[GetGoroutineID()]) == 0 {
		delete(cdb, GetGoroutineID())
	}
}

func connectDB() {
	dbc = make(MG)
	cdb = make(map[uint64]MG)
	dbConf := dbConf()
	if len(dbConf.(map[interface{}]interface{})) > 1 {
		for k, v := range dbConf.(map[interface{}]interface{}) {
			openDB(albedo.MakeString(k), v.(map[interface{}]interface{}))
		}
	} else {
		openDB("db", dbConf.(map[interface{}]interface{}))
	}
}

func openDB(name string, options map[interface{}]interface{}) {
	ms := make(MS)
	for k, v := range options {
		ms[albedo.MakeString(k)] = albedo.MakeString(v)
	}
	if ms["charset"] == "" {
		ms["charset"] = "utf8"
	}
	db, err := gorm.Open(ms["datatype"], ms["username"]+":"+ms["password"]+"@tcp("+ms["hostname"]+")/"+ms["database"]+"?charset="+ms["charset"])
	errInfo := ""
	if err != nil {
		errInfo = err.Error()
	}
	Assert(err == nil, errInfo)
	//全局禁用表复数
	db.SingularTable(true)
	//openDB.LogMode(true)
	//openDB.SetLogger(newDBLog(name))
	dbc[name] = db
}

func dbConf() interface{} {
	dbConf := Config("db")
	Assert(dbConf != nil, "db config is not set")
	if len(dbConf.(map[interface{}]interface{})) > 1 {
		for k, _ := range dbConf.(map[interface{}]interface{}) {
			dbs = append(dbs, albedo.MakeString(k))
		}
	}
	return dbConf
}
