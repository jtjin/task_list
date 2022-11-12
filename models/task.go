package models

type Task struct {
	Id     int    `gorm:"column:id;primary_key;AUTO_INCREMENT;type:int(11)"`
	Name   string `gorm:"column:name;NOT NULL;type:varchar(255)"`           // 名稱
	Status int    `gorm:"column:status;default:0;NOT NULL;type:tinyint(1)"` // 狀態 (0: 未完成, 1: 已完成)
}
