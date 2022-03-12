package conf

import (
	"fin"
	"os"
	"path"
)

func Run(engine *fin.Engine) {
	dir, _ := os.Getwd()
	engine.SetConfig(path.Join(dir, "/example/config/config.yml"))
	engine.SetMessage(path.Join(dir, "/example/config/message.yml"))
}
