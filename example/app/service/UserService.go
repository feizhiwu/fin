package service

import (
	"encoding/json"
	"fin"
	"fin/example/app/dao"
	"github.com/feizhiwu/gs/albedo"
)

type UserService struct {
	UD *dao.UserDao
}

func (s *UserService) Add(data map[string]interface{}) {
	s.UD = dao.User()
	params := fin.CopyParams([]string{"name", "password"}, data)
	json.Unmarshal(albedo.MakeJson(params), &s.UD.User)
	s.UD.User.Password = fin.EncryptPass(s.UD.User.Password)
	s.UD.Add()
}

func (s *UserService) GetInfo(id uint) {
	s.UD = dao.User()
	s.UD.User.Id = id
	s.UD.GetOne()
}

func (s *UserService) Update(data map[string]interface{}) {
	s.UD = dao.User()
	params := fin.CopyParams([]string{"id", "name", "password"}, data)
	s.UD.Update(params)
	panic(1111)
}

func (s *UserService) Delete(id uint) {
	s.UD = dao.User()
	s.UD.User.Id = id
	s.UD.Delete()
}

func (s *UserService) GetList(data map[string]interface{}) {
	s.UD = dao.User()
	s.UD.GetAll(data)
}
