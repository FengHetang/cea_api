/**
 * @Author: 戈亓
 * @Description:
 * @File:  UnitsServer
 * @Version: 1.0.0
 * @Date: 2022/7/28 16:18
 */

package UnitsServer

import "cea_api/models"

type Unit struct {
	Unit string
	Mark string
}

func (u *Unit) AddUnit() (res string) {
	return models.AddUnit(u.Unit, u.Mark)
}

type AddDepart struct {
	Department string
	UnitMark   string
}

func (d *AddDepart) AddDepartment() (res string) {
	return models.AddDepaerment(d.Department, d.UnitMark)
}
