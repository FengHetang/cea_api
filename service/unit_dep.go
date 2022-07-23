/**
 * @Author: 戈亓
 * @Description:
 * @File:  unit_dep
 * @Version: 1.0.0
 * @Date: 2022/7/24 0:48
 */

package service

import "cea_api/models"

func GerDepart(unitmark string) (departlist []models.Department) {
	departlist = []models.Department{}
	models.DB.Where("UnitMark = ? ", unitmark).Find(&departlist)
	return departlist
}
