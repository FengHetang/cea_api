/**
 * @Author: 戈亓
 * @Description:
 * @File:  UserServer
 * @Version: 1.0.0
 * @Date: 2022/7/24 20:33
 */

package UserServer

import "cea_api/models"

func UpdateUserPwd(oldpwd, newpwd, token string) (res string) {
	return models.UserUpdatePwd(oldpwd, newpwd, token)
}

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
