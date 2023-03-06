package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/models"
)

// GetUserList
// @Tags         用户
// @Success      200  {string}  json{"code","message"}
// @Router       /user [get]
func GetUserList(c *gin.Context) {
	user_list := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": user_list,
	})
}
