package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"project/models"
	"project/utils"
	"strconv"
)

// GetUserList
// @Summary      获取用户
// @Tags         用户
// @Success      200  {string}  json{"code","message"}
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	user_list := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": user_list,
	})
}

// CreateUser
// @Summary      新增用户
// @Tags         用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success      200  {string}  json{"code","message"}
// @Router       /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.User{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	salt := fmt.Sprint("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

// DeleteUser
// @Summary      删除用户
// @Tags         用户
// @param id query string false "用户id"
// @Success      200  {string}  json{"code","message"}
// @Router       /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary      更新用户
// @Tags         用户
// @param id formData string false "用户id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success      200  {string}  json{"code","message"}
// @Router       /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	fmt.Println("upate: ", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
		})
	} else {
		user.ID = uint(id)
		models.UpdateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "更新成功",
		})
	}
}
