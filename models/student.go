package models

type Student struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Course string `json:"course"`
}
