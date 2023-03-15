package models

import (
	"fmt"
	"gorm.io/gorm"
	"project/utils"
	"time"
)

type User struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	ClientIp      string
	Identity      string
	ClientPort    string
	Salt          string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogOutTime    *time.Time `gorm:"column:log_out_time" json:"login_out_time"`
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
	return utils.DB.Model(&user).Updates(User{Name: user.Name, PassWord: user.PassWord, Email: user.Email, Phone: user.Phone})
}
