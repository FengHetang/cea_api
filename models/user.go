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

type LoginUserRes struct {
	Realname   string
	Unit       string
	Department string
	Usertype   string
}

func GetUserData(name, unit, department string) (userid, realname, usertype string) {
	var user User
	DB.Where("name = ? and unit = ? and department = ?", name, unit, department).Find(&user)
	if user.ID > 0 {
		userid := user.Userid
		realname := user.Realname
		usertype := user.Usertype
		return userid, realname, usertype
	} else {
		return
	}
}

func LoginCheck(name, password, unit, department string) (code int) {
	var user User
	DB.Where("name = ? and unit = ? and department = ?", name, unit, department).Find(&user)
	if user.ID > 0 {
		if user.Password != password {
			return 20004
		} else {
			return 200
		}
	} else {
		return 20002
	}
}

//
func ValOldPwd(oldpwd, token string) (res bool) {
	var user User
	tokenvaldata, _ := jwt.ParseToken(token)
	UserId := tokenvaldata.UserId
	DB.Where("userid = ? ", UserId).Find(&user)
	if user.Password == oldpwd {
		return true
	} else {
		return false
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
