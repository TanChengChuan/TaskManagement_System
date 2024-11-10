package models

import (
	"errors"
	"gorm.io/gorm"
)

var DSN = "root:hsw20050529@tcp(127.0.0.1:3306)/taskmanagement?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func CheckUserCredentials(username string, password string) (bool, error) { //检查用户名是否存在
	var user User
	result := DB.First(&user, "username = ?", username)
	if result.Error != nil {
		return false, result.Error
	} //查询出错
	if result.RowsAffected == 0 {
		return false, errors.New("user or password is invalid") //不管是用户不存在还是密码不匹配，都整一样的。因为其他程序也这么做的（防止别人猜吧）
	}
	if password != user.Password { //后面改为哈希加密
		return false, errors.New("user or password is invalid")
	}
	return true, nil
}
