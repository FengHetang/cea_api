/**
 * @Author: 戈亓
 * @Description:
 * @File:  userRouters
 * @Version: 1.0.0
 * @Date: 2022/7/17 21:46
 */

package router

import (
	"cea_api/controller"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	UserRouter := r.Group("/user")
	{
		UserRouter.GET("/get_vercode", controller.GetVerCode)
		UserRouter.GET("/get_unit", controller.GetUnit)
		UserRouter.GET("/get_department", controller.GetDepart)
		UserRouter.GET("/val_vercode", controller.ValVerCode)
		UserRouter.GET("/login", controller.Login)
		UserRouter.GET("/login1", controller.ValToken)
	}
}
