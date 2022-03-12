package controller

import (
	"fin"
)

type UserController struct {
	*fin.Display
	params fin.MI
}

func User(c *fin.Context) {
	s := &UserController{
		c.Display(),
		c.GetParams(),
	}
	s.Get(s.list)
	s.Get(s.info)
	s.Post(s.add)
	s.Put(s.update)
	s.Delete(s.delete)
	s.Run()
}

func (s *UserController) list() {
	s.Show([]fin.MS{
		{
			"name": "tout",
			"job":  "no",
		},
		{
			"name": "phia",
			"job":  "no",
		},
	})
}

func (s *UserController) info() {
	s.Show(fin.MS{
		"name": "tout",
		"job":  "no",
	})
}

func (s *UserController) add() {
	s.Show(fin.StatusOK)
}

func (s *UserController) update() {
	s.Show(fin.StatusOK)
}

func (s *UserController) delete() {
	s.Show(fin.StatusOK)
}
