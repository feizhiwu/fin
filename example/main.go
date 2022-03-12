package main

import (
	"fin"
	"fin/example/config/conf"
	"fin/example/config/route"
)

func main() {
	engine := fin.New()
	conf.Run(engine)
	route.Run(engine)
	engine.Run(":8080")
}
