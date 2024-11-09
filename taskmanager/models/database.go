package models

import "gorm.io/gorm"

var DSN = "root:hsw20050529@tcp(127.0.0.1:3306)/taskmanagement?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB
