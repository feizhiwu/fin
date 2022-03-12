package controller

import (
	"fin"
)

func Index(c *fin.Context) {
	c.Display().Show(fin.StatusOK)
}
