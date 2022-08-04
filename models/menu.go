/**
 * @Author: 戈亓
 * @Description:
 * @File:  menu
 * @Version: 1.0.0
 * @Date: 2022/8/4 21:15
 */

package models

type Menu struct {
	Id       int
	MenuName string
	OrderNum string
	Menuicon string
}

func (Menu) TableName() string {
	return "menu"
}

type MenuItem struct {
	Id       int
	MenuId   int
	ItemName string
	OrderNum string
	Url      string
	Itemicon string
}

func (MenuItem) TableName() string {
	return "menu_item"
}

type UserType struct {
	Id       int
	UserType string
}
