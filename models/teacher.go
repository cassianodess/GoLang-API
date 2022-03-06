package models

type Teacher struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}
