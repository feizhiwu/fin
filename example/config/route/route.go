package route

import (
	"fin"
	"fin/example/app/controller"
	"fin/example/app/middleware"
)

type route struct {
	Engine *fin.Engine
}

// Run 路由路口
func Run(engine *fin.Engine) {
	r := route{engine}
	r.Engine.Use(middleware.Cors)
	r.Engine.Any("", controller.Index)
	r.v1()
}

func (r *route) v1() {
	v1 := r.Engine.Group("v1")
	{
		v1.Any("/user", controller.User)
	}
}
