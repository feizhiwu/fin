package dao

import (
	"fin"
	"fin/example/app/model"
	"github.com/feizhiwu/gs/kokomi"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	DB       *gorm.DB
	User     model.User
	UserList model.UserList
}

func User() *UserDao {
	return &UserDao{
		DB: fin.BeginDB("main_db"),
	}
}

func (d *UserDao) Add() {
	d.DB.Create(&d.User)
}

func (d *UserDao) Update(data map[string]interface{}) {
	d.DB.Table("user").Where("id  = ?", data["id"]).Updates(data)
}

func (d *UserDao) GetOne() {
	d.DB.Where("id  = ?", d.User.Id).First(&d.User)
}

func (d *UserDao) Delete() {
	d.DB.Table("user").Delete(&d.User)
}

func (d *UserDao) GetAll(data map[string]interface{}) {
	db := d.DB.Model(model.User{})
	query := kokomi.NewQuery(&db, data)
	query.Like("name") //如果传参data["name"]，则进行like匹配查询
	query.List(&d.UserList.List).Pages(&d.UserList.Pages)
}
