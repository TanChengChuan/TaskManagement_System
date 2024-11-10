package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Task struct { //记得加数据校验
	ID          uint      `gorm:"primary_key" json:"TaskID" bind:"required" mapstructure:"TaskID"`
	Title       string    ` json:"Title" bind:"required" mapstructure:"Title"`
	Description string    ` json:"Description" bind:"required" mapstructure:"Description"`
	Status      int       `json:"Status"  bind:"required" mapstructure:"Status"` //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time `json:"CreatedAt"  bind:"required" mapstructure:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"  bind:"required" mapstructure:"UpdatedAt"`
	Owners      []User    `gorm:"many2many:task_user" bind:"required" mapstructure:"users"`
}

type TaskBatch struct {
	Tasks []Task `json:"tasks" bind:"required"`
}
