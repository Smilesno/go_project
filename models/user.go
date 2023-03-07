package models

import (
	"fmt"
	"gorm.io/gorm"
	"project/utils"
)

type User struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	ClientIp      string
	Identity      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogOutTime    uint64
	Islogout      bool
	DeviceInfo    string
}

func (table *User) TableName() string {
	return "user"
}

func GetUserList() []*User {
	data := make([]*User, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Model(&user).Updates(User{Name: user.Name, PassWord: user.PassWord})
}
