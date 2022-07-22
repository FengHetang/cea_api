/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:25
 */

package controller

import (
	"cea_api/models"
	"cea_api/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userLiset := []models.User{}
	models.DB.Find(&userLiset)
	c.JSON(200, gin.H{
		"resule": userLiset,
	})
}
func Ver_code(c *gin.Context) {
	id, b64s := service.GetCaptcha()
	c.JSON(200, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}
