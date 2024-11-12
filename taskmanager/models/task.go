package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Task struct { //记得加数据校验
	ID          uint      `gorm:"primary_key"  mapstructure:"TaskID"`
	Title       string    ` json:"Title" bind:"required"  mapstructure:"Title"`
	Description string    ` json:"Description" bind:"required"  mapstructure:"Description"`
	Status      int       `json:"Status"  bind:"required"  mapstructure:"Status"` //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time `json:"CreatedAt"  mapstructure:"CreatedAt" `
	UpdatedAt   time.Time `json:"UpdatedAt"  mapstructure:"UpdatedAt" `
	OwnerID     uint      `gorm:"index"`
	//Owner       User      `gorm:"foreignKey:OwnerID" json:"Owner"  bind:"required" `
}

type TaskBatch struct {
	Tasks []Task `json:"tasks"  bind:"required"`
}
