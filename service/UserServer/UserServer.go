/**
 * @Author: 戈亓
 * @Description:
 * @File:  UserServer
 * @Version: 1.0.0
 * @Date: 2022/7/24 20:33
 */

package UserServer

import (
	"cea_api/models"
	"cea_api/pkg/jwt"
)

type User struct {
	Userid     string
	Realname   string
	Name       string
	Password   string
	Unit       string
	Department string
	Usertype   string
}

func (u *User) AddUser() (res string) {
	return models.AddUser(u.Usertype, u.Unit, u.Realname, u.Department, u.Userid, u.Name, u.Password)
}

func ValOldPwd(oldpwd, token string) (res bool) {
	return models.ValOldPwd(oldpwd, token)
}

func UpdateUserPwd(newpwd, token string) (res bool) {
	return models.UserUpdatePwd(newpwd, token)
}

type UserLogin struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
}

type LoginRes struct {
	Name       string
	Realname   string `json:"realname"`
	Usertype   string `json:"usertype"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
	Token      string `json:"token"`
}

func (u *UserLogin) LoginCheck() (code int) {
	return models.LoginCheck(u.Name, u.Password, u.Unit, u.Department)
}

func (l *LoginRes) CreatToken() (res interface{}) {
	userid, realname, usertype := models.GetUserData(l.Name, l.Unit, l.Department)
	token_ := jwt.TokenData{Realname: realname, UserType: usertype, Unit: l.Unit, Department: l.Department, UserId: userid}
	token, _ := token_.GenToken(userid, realname, usertype, l.Unit, l.Department)
	result := LoginRes{Name: l.Name, Unit: l.Unit, Department: l.Department, Realname: realname, Usertype: usertype, Token: token}
	res = result
	return res
}
