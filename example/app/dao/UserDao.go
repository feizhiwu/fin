package dao

import (
	"fin/example/app/model"
	"fin/example/config/conf"
	"github.com/feizhiwu/gs/kokomi"
)

type UserDao struct {
	User     model.User
	UserList model.UserList
}

func (d *UserDao) Add() {
	conf.MainDB.Create(&d.User)
}

func (d *UserDao) Update(data map[string]interface{}) {
	conf.MainDB.Table("user").Where("id  = ?", data["id"]).Updates(data)
}

func (d *UserDao) GetOne() {
	conf.MainDB.Where("id  = ?", d.User.Id).First(&d.User)
}

func (d *UserDao) Delete() {
	conf.MainDB.Table("user").Delete(&d.User)
}

func (d *UserDao) GetAll(data map[string]interface{}) {
	db := conf.MainDB.Model(model.User{})
	query := kokomi.NewQuery(&db, data)
	query.Like("name") //如果传参data["name"]，则进行like匹配查询
	query.List(&d.UserList.List).Pages(&d.UserList.Pages)
}