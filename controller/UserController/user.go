/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:25
 */

package UserController

import (
	"cea_api/pkg/app"
	"cea_api/pkg/e"
	"cea_api/pkg/jwt"
	"cea_api/service"
	"cea_api/service/UserServer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetVerCode 获取验证码
func GetVerCode(c *gin.Context) {
	app := app.Gin{c}
	res := service.GetCaptcha()
	app.Response(200, e.Success, res)
}

// ValVerCode 验证验证码
func ValVerCode(c *gin.Context) {
	app := app.Gin{c}
	id := c.Query("codeid")
	ret_captcha := c.Query("vercode")
	res := service.VerityCaptcha(id, ret_captcha)
	if res == false {
		app.Response(http.StatusOK, e.Valerror, nil)
	} else {
		app.Response(http.StatusOK, e.Valright, nil)
	}
}

// Login 登录
func Login(c *gin.Context) {
	app := app.Gin{c}
	username := c.Query("username")
	password := c.Query("password")
	unit := c.Query("unit")
	department := c.Query("department")
	user := UserServer.UserLogin{username, password, unit, department}
	code := user.LoginCheck()
	switch code {
	case 200:
		result := UserServer.LoginRes{Name: username, Department: department, Unit: unit}
		res := result.CreatToken()
		app.Response(http.StatusOK, e.Success, res)
	case 20004:
		app.Response(http.StatusOK, e.UserPasswordError, nil)
	case 20002:
		app.Response(http.StatusOK, e.UserNoExits, nil)
	default:
		app.Response(http.StatusOK, e.Error, nil)
	}
}

// ValToken 验证token
func ValToken(c *gin.Context) {
	app := app.Gin{c}
	token := c.Query("token")
	tokenvaldata, _ := jwt.ParseToken(token)
	fmt.Println(tokenvaldata)
	if tokenvaldata != nil {
		app.Response(http.StatusOK, e.Success, tokenvaldata)
	} else {
		app.Response(http.StatusOK, e.CheckTokenTimeout, nil)
	}
}

// 旧密码验证
func ValOldPwd(c *gin.Context) {
	app := app.Gin{c}
	oldpwd := c.Query("oldpwd")
	token := c.Request.Header.Get("token")
	res := UserServer.ValOldPwd(oldpwd, token)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.OldPasswordError, nil)
	}
}

// UserUpdatePwd 修改密码
func UserUpdatePwd(c *gin.Context) {
	app := app.Gin{c}
	NewPwd := c.Query("newpwd")
	token := c.Request.Header.Get("token")
	res := UserServer.UpdateUserPwd(NewPwd, token)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.EditOldPasswordError, nil)
	}
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
