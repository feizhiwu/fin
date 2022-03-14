package controller

import (
	"fin"
)

type IndexController struct {
	*fin.Display
}

func Index(c *fin.Context) {
	s := &IndexController{
		c.NewDisplay(),
	}
	s.Show(10000)
}
