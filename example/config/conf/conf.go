package conf

import (
	"fin"
	"github.com/jinzhu/gorm"
	"os"
	"path"
)

var (
	MainDB   *gorm.DB
	ClientDB *gorm.DB
)

func Run(engine *fin.Engine) {
	dir, _ := os.Getwd()
	engine.SetConfig(path.Join(dir, "/example/config/config.yml"))
	engine.SetMessage(path.Join(dir, "/example/config/message.yml"))
	db := engine.ConnectDB()
	MainDB = db["main_db"]
	ClientDB = db["client_db"]
}
