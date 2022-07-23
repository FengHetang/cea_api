/**
 * @Author: 戈亓
 * @Description:
 * @File:  unit
 * @Version: 1.0.0
 * @Date: 2022/7/24 0:02
 */

package models

type Unit struct {
	Unit string `json:"unit"`
	Mark string `json:"mark"`
}

func (Unit) TableName() string {
	return "Unit"
}

type Department struct {
	Department string `json:"department"`
	SerNum     string `json:"serNum"`
	UnitMark   string `json:"unitMark"`
}

func (Department) TableName() string {
	return "Department"
}