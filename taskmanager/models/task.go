package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Task struct { //记得加数据校验
	ID          uint      `gorm:"primary_key" yaml:"TaskID" bind:"required"`
	Title       string    ` yaml:"Title" bind:"required"`
	Description string    ` yaml:"Description" bind:"required"`
	Status      int       `yaml:"Status"  bind:"required"` //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time `yaml:"CreatedAt"  bind:"required"`
	UpdatedAt   time.Time `yaml:"UpdatedAt"  bind:"required"`
	Owners      []User    `gorm:"many2many:task_user" bind:"required"`
}

type TaskBatch struct {
	Tasks []Task `yaml:"tasks" bind:"required"`
}
