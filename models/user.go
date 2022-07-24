/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:21
 */

package models

import (
	"cea_api/pkg/jwt"
)

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

func LoginCheck(name, password, unit, department string) (realneme, usertype, userid string) {
	var user User
	DB.Where("name = ? and password = ? and unit = ? and department = ?", name, password, unit, department).First(&user)
	if user.ID > 0 {
		realneme = user.Realname
		usertype = user.Usertype
		userid = user.Userid
		return realneme, usertype, userid
	} else {
		realneme = ""
		usertype = ""
		userid = ""
		return realneme, usertype, userid
	}
}
func UserUpdatePwd(oldpwd, newpwd, token string) (res string) {
	userdata, _ := jwt.ParseToken(token)
	var user User
	DB.Where("userid = ?", userdata.UserId).First(&user)
	if user.Password == oldpwd {
		DB.Model(user).Where("userid = ?", userdata.UserId).Update("password", newpwd)
		return "密码更新成功！"
	} else {
		return "旧密码输入错误！"
	}
}
