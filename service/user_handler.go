package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/models"
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
	user.PassWord = password
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
// @Success      200  {string}  json{"code","message"}
// @Router       /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.ID = uint(id)
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}
