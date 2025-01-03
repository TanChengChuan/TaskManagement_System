package models

// 加入数据库
type User struct {
	ID       uint   `gorm:"primaryKey" json:"UserID"  mapstructure:"UserID"` //也是自动添加的
	Username string `json:"username" bind:"required" mapstructure:"username" `
	Password string `json:"password" bind:"required" mapstructure:"password"` //记得保护密码
	Tasks    []Task `gorm:"foreignKey:OwnerID;references:ID" `
}
