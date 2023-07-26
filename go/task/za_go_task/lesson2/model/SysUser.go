package model

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Name string
	Age  int
}
