package models

type Task struct {
	Id     int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"column:name;NOT NULL"`             // 名稱
	Status int    `gorm:"column:status;default:0;NOT NULL"` // 狀態 (0: 未完成, 1: 已完成)
}
