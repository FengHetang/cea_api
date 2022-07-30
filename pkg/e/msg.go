/**
 * @Author: 戈亓
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2022/7/30 14:44
 */

package e

var magflags = map[int]string{
	Success:       "OK",
	Error:         "FAIL",
	Invalidparams: "请求参数错误",

	CheckTokenFail:    "token验证失败",
	CheckTokenTimeout: "token 验证超时",
	CreateTokenError:  "token创建失败",
	TokenNull:         "无token数据",

	Valerror: "验证码错误",
	Valright: "验证码正确",

	UserExits:         "用户已存在",
	UserNoExits:       "用户不存在",
	UserCreate:        "用户创建失败",
	UserPasswordError: "密码错误",
	OldPasswordError:  "旧密码错误",
}

func GetMsg(code int) string {
	msg, ok := magflags[code]
	if ok {
		return msg
	}
	return magflags[Error]
}
