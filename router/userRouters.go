/**
 * @Author: 戈亓
 * @Description: Routing branch
 * @File:  userRouters
 * @Version: 1.0.0
 * @Date: 2022/7/17 21:46
 */

package router

import (
	"cea_api/controller/UserController"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	UserRouter := r.Group("/user")
	{
		UserRouter.GET("/get_vercode", UserController.GetVerCode)   // 获取验证码
		UserRouter.GET("/val_vercode", UserController.ValVerCode)   // 验证 验证码
		UserRouter.GET("/login", UserController.Login)              // 登录
		UserRouter.GET("/ValToken", UserController.ValToken)        //验证token
		UserRouter.POST("/updatepwd", UserController.UserUpdatePwd) //修改密码
		UserRouter.POST("/adduser", UserController.UserAdd)
	}
}
