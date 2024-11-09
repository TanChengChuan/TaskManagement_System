package models

// 加入数据库
type User struct {
	ID       uint   `yaml:"UserID" mapstructure:"UserID"` //也是自动添加的
	Username string `yaml:"username" bind:"required" mapstructure:"username" `
	Password string `yaml:"password" bind:"required" mapstructure:"password"` //记得保护密码
	Tasks    []Task `gorm:"many2many:task_user" yaml:"tasks" bind:"required" `
}
