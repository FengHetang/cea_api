/**
 * @Author: 戈亓
 * @Description:
 * @File:  units
 * @Version: 1.0.0
 * @Date: 2022/7/28 15:18
 */

package UnitsController

import (
	"cea_api/models"
	"cea_api/service"
	"cea_api/service/UnitsServer"
	"github.com/gin-gonic/gin"
)

func GetDepart(c *gin.Context) {
	unitmark := c.Query("unit_mark")
	departlist := service.GerDepart(unitmark)
	c.JSON(200, gin.H{
		"data": departlist,
	})
}

func GetUnit(c *gin.Context) {
	unitList := []models.Unit{}
	models.DB.Find(&unitList)
	c.JSON(200, gin.H{
		"data": unitList,
	})
}

func AddUnit(c *gin.Context) {
	unit := c.Query("unitname")
	mark := c.Query("mark")
	addserver := &UnitsServer.Unit{unit, mark}
	res := addserver.AddUnit()
	c.JSON(200, gin.H{
		"res": res,
	})
}

func AddDeartment(c *gin.Context) {
	unitmark := c.Query("unitmark")
	department := c.Query("deapartment")
	addserver := &UnitsServer.AddDepart{UnitMark: unitmark, Department: department}
	res := addserver.AddDepartment()
	c.JSON(200, gin.H{
		"res": res,
	})
}
