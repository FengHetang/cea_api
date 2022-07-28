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
	"fmt"
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

func LoginCheck(name, password, unit, department string) (realneme, usertype, userid, msg string) {
	var user User
	DB.Where("name = ? and unit = ? and department = ?", name, unit, department).First(&user)
	if user.ID > 0 {
		if user.Password != password {
			msg = "密码错误"
			return realneme, usertype, userid, msg
		} else {
			realneme = user.Realname
			usertype = user.Usertype
			userid = user.Userid
			return realneme, usertype, userid, msg
		}
	} else {
		msg = "查无此人"
		return realneme, usertype, userid, msg
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

func AddUser(usertype, unit, realname, deaprtment, userid, username, password string) (res string) {
	var user User
	DB.Where("userid = ?", userid).First(&user)
	if user.ID > 0 {
		return "该用户已存在！"
	} else {
		result := DB.Create(&User{Userid: userid, Name: username, Password: password, Realname: realname, Usertype: usertype, Unit: unit, Department: deaprtment})
		if result.Error != nil {
			fmt.Println(result.Error)
			return "新增用户失败"
		}
		return "新增用户成功！"
	}
}
