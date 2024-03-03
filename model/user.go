package model

import (

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email    string
	PassWord string
	NickName string
	Status   string
	Avatar   string
	Money    string
}

const (
	PassWordCost = 12 // 密码难度 
	Active string = "active" // 用户状态
)

// 设置密码  
func (user *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost) // bcrypt.GenerateFromPassword
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
} 