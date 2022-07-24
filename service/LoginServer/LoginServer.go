/**
 * @Author: 戈亓
 * @Description:
 * @File:  LoginServer
 * @Version: 1.0.0
 * @Date: 2022/7/24 15:14
 */

package LoginServer

import "cea_api/models"

type UserLogin struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
}

func (u *UserLogin) LoginCheck() (realneme, usertype, userid, res string) {
	return models.LoginCheck(u.Name, u.Password, u.Unit, u.Department)
}
