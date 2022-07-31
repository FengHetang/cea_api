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
	"cea_api/pkg/app"
	"cea_api/pkg/e"
	"cea_api/service/UnitsServer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDepart(c *gin.Context) {
	app := app.Gin{c}
	unitmark := c.Query("unit_mark")
	departlist := []models.Department{}
	result := models.DB.Where("unit_mark = ? ", unitmark).Find(&departlist)
	if result.Error == nil {
		app.Response(http.StatusOK, e.Success, departlist)
	} else {
		app.Response(http.StatusOK, e.DepartGetError, nil)
	}
}

func GetUnit(c *gin.Context) {
	app := app.Gin{c}
	unitList := []models.Unit{}
	result := models.DB.Find(&unitList)
	if result.Error == nil {
		app.Response(http.StatusOK, e.Success, unitList)
	} else {
		app.Response(http.StatusOK, e.UnitGetError, nil)
	}

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
