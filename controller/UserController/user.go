/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:25
 */

package UserController

import (
	"cea_api/pkg/jwt"
	"cea_api/service"
	"cea_api/service/LoginServer"
	"cea_api/service/UserServer"
	"github.com/gin-gonic/gin"
)

func GetVerCode(c *gin.Context) {
	id, b64s := service.GetCaptcha()
	c.JSON(200, gin.H{
		"codeid":  id,
		"b64data": b64s,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	unit := c.Query("unit")
	department := c.Query("department")
	user := LoginServer.UserLogin{username, password, unit, department}
	realneme, usertype, userid, res := user.LoginCheck()
	if res != "" {
		c.JSON(200, gin.H{
			"msg": res,
		})
	} else {
		token_ := jwt.TokenData{Realname: realneme, UserType: usertype, Unit: unit, Department: department, UserId: userid}
		token, _ := token_.GenToken(userid, realneme, usertype, unit, department)
		c.JSON(200, gin.H{
			"realneme": realneme,
			"usertype": usertype,
			"token":    token,
		})
	}
}
func ValToken(c *gin.Context) {
	token := c.Query("token")
	tokenvaldata, _ := jwt.ParseToken(token)
	if tokenvaldata != nil {
		c.JSON(200, gin.H{
			"data": tokenvaldata,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "token已过期",
		})
	}
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

func UserUpdatePwd(c *gin.Context) {
	NewPwd := c.Query("newpwd")
	OldPwd := c.Query("oldpwd")
	token := c.Request.Header.Get("token")
	res := UserServer.UpdateUserPwd(OldPwd, NewPwd, token)
	c.JSON(200, gin.H{
		"res": res,
	})
}

func UserAdd(c *gin.Context) {
	userid := c.Query("userid")
	unit := c.Query("unit")
	deaprtment := c.Query("department")
	realname := c.Query("realname")
	username := c.Query("username")
	password := c.Query("password")
	usertype := c.Query("userptype")
	token := c.Request.Header.Get("token")
	userdata, _ := jwt.ParseToken(token)
	if userdata == nil {
		c.JSON(200, gin.H{
			"res": "token已过期",
		})
	} else {
		if userdata.UserType == "管理员" && userdata.Unit != unit {
			c.JSON(200, gin.H{
				"res": "仅支持创建本单位新用户",
			})
		} else {
			addserver := &UserServer.User{userid, realname, username, password, unit, deaprtment, usertype}
			res := addserver.AddUser()
			c.JSON(200, gin.H{
				"res": res,
			})
		}
	}

}
