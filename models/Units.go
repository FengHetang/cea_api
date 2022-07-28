/**
 * @Author: 戈亓
 * @Description:
 * @File:  unit
 * @Version: 1.0.0
 * @Date: 2022/7/24 0:02
 */

package models

import "fmt"

type Unit struct {
	Id   int    `json:"id"`
	Unit string `json:"unit"`
	Mark string `json:"mark"`
}

func (Unit) TableName() string {
	return "Unit"
}

func (Department) TableName() string {
	return "Department"
}

func AddUnit(unit, mark string) (res string) {
	var u Unit
	DB.Where("Unit=?", unit).First(&u)
	fmt.Println(u)
	if u.Id > 0 {
		return "该单位已存在"
	} else {
		result := DB.Create(&Unit{Unit: unit, Mark: mark})
		if result.Error != nil {
			return "新增单位失败"
		}
		return "新增单位成功"
	}
}

type Department struct {
	ID         int    `json:"id"`
	Department string `json:"department"`
	SerNum     string `json:"serNum"`
	UnitMark   string `json:"unitMark"`
}

func AddDepaerment(department, unitmark string) (res string) {
	var d Department
	DB.Where("Department = ? and Unitmark = ?", department, unitmark).Find(&d)
	if d.ID > 0 {
		return "该部门已存在"
	} else {
		departmentList := []Department{}
		result := DB.Where("UnitMark=?", unitmark).Find(&departmentList)
		fmt.Println(result)
	}
	return "1111"
}
