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
