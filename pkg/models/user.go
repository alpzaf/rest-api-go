package models

type User struct {
	Id int `json:"id" gorm: "primaryKey"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Desc string `json:"desc"`
}