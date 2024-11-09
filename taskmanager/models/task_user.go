package models

import "database/sql"

// TaskUser 多对多关联定义.
type TaskUser struct {
	Task    Task          `gorm:"foreignKey:TaskID"`
	TaskID  sql.NullInt64 `yaml:"TaskID"`
	User    User          `gorm:"foreignKey:OwnerID"`
	OwnerID sql.NullInt64 `yaml:"UserID"`
}
