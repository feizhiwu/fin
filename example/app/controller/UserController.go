package controller

import (
	"fin"
	"fin/example/app/service"
	"github.com/feizhiwu/gs/albedo"
)

type UserController struct {
	*fin.Display
	params fin.MI
	us     *service.UserService
}

func User(c *fin.Context) {
	s := &UserController{
		c.Display(),
		c.GetParams(),
		new(service.UserService),
	}
	s.Get(s.list)
	s.Get(s.info)
	s.Post(s.add)
	s.Put(s.update)
	s.Delete(s.delete)
	s.Run()
}

func (s *UserController) add() {
	val := map[int]string{
		20001: "name",
		20002: "password",
	}
	s.Validate(val, s.params)
	s.us.Add(s.params)
	data := map[string]uint{
		"id": s.us.UD.User.Id,
	}
	s.Show(data)
}

func (s *UserController) list() {
	val := map[int]string{
		80007: "page",
	}
	s.Validate(val, s.params)
	s.us.GetList(s.params)
	s.Show(s.us.UD.UserList)
}

func (s *UserController) info() {
	s.HasKey(s.params)
	s.us.GetInfo(albedo.MakeUint(s.params["id"]))
	s.Show(s.us.UD.User)
}

func (s *UserController) update() {
	s.HasKey(s.params)
	s.us.Update(s.params)
	s.Show(fin.StatusOK)
}

func (s *UserController) delete() {
	s.HasKey(s.params)
	s.us.Delete(albedo.MakeUint(s.params["id"]))
	s.Show(fin.StatusOK)
}
