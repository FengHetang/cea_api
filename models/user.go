/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:21
 */

package models

type User struct {
	Id       int
	Realname string
	Name     string
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
