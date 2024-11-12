package models

// TaskUser 多对多关联定义.  -》改为一对多 	多对多太复杂了
/*type TaskUser struct {
	Task    Task          `gorm:"foreignKey:TaskID"`
	TaskID  sql.NullInt64 `json:"TaskID"`
	User    User          `gorm:"foreignKey:OwnerID"`
	OwnerID sql.NullInt64 `json:"UserID"`
}
*/
