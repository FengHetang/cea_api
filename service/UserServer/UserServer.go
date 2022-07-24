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
