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

func GetVerCode(c *gin.Context) {
	id, b64s := service.GetCaptcha()
	c.JSON(200, gin.H{
		"codeid":  id,
		"b64data": b64s,
	})
}

func GetUnit(c *gin.Context) {
	unitList := []models.Unit{}
	models.DB.Find(&unitList)
	c.JSON(200, gin.H{
		"data": unitList,
	})
}
func GetDepart(c *gin.Context) {
	unitmark := c.Query("unit_mark")
	departlist := service.GerDepart(unitmark)
	c.JSON(200, gin.H{
		"data": departlist,
	})
}

func Login(c *gin.Context) {
	//unit := c.Query("unit")
	//department := c.Query("department")
	//username := c.Query("username")
	//password := c.Query("password")
	userLiset := []models.User{}
	models.DB.Find(&userLiset)
	c.JSON(200, gin.H{
		"resule": userLiset,
	})
}

func ValVerCode(c *gin.Context) {
	id := c.Query("codeid")
	ret_captcha := c.Query("vercode")
	res := service.VerityCaptcha(id, ret_captcha)
	if res == false {
		c.JSON(200, gin.H{
			"res": "验证码错误",
		})
	} else {
		c.JSON(200, gin.H{
			"res": "验证码正确",
		})
	}
}
