package conf

import (
	"fin"
	"os"
	"path"
	"time"
)

func Run(engine *fin.Engine) {
	dir, _ := os.Getwd()
	engine.SetConfig(path.Join(dir, "/example/config/config.yml"))
	engine.SetMessage(path.Join(dir, "/example/config/message.yml"))
	engine.SetLog(path.Join(dir, "/example/data/runtime/log/"+time.Now().Format("200601")))
}
