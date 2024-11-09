package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Task struct { //记得加数据校验
	ID          uint      `gorm:"primary_key" yaml:"TaskID" bind:"required" mapstructure:"TaskID"`
	Title       string    ` yaml:"Title" bind:"required" mapstructure:"Title"`
	Description string    ` yaml:"Description" bind:"required" mapstructure:"Description"`
	Status      int       `yaml:"Status"  bind:"required" mapstructure:"Status"` //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time `yaml:"CreatedAt"  bind:"required" mapstructure:"CreatedAt"`
	UpdatedAt   time.Time `yaml:"UpdatedAt"  bind:"required" mapstructure:"UpdatedAt"`
	Owners      []User    `gorm:"many2many:task_user" bind:"required" mapstructure:"users"`
}

type TaskBatch struct {
	Tasks []Task `yaml:"tasks" bind:"required"`
}
