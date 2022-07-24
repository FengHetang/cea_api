/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:21
 */

package models

type User struct {
	ID         int
	Userid     string
	Realname   string `json:"realname"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
	Usertype   string `json:"usertype"`
}

func (User) TableName() string {
	return "User"
}

func LoginCheck(name, password, unit, department string) (realneme, usertype string) {
	var user User
	DB.Where("name = ? and password = ? and unit = ? and department = ?", name, password, unit, department).First(&user)
	if user.ID > 0 {
		realneme = user.Realname
		usertype = user.Usertype
		return realneme, usertype
	} else {
		realneme = ""
		usertype = ""
		return realneme, usertype
	}
}
