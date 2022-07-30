/**
 * @Author: 戈亓
 * @Description:
 * @File:  code
 * @Version: 1.0.0
 * @Date: 2022/7/30 13:21
 */

package e

const (
	Success       = 200
	Error         = 500
	Invalidparams = 400 //参数无效

	CheckTokenFail    = 10001 // token验证失败
	CheckTokenTimeout = 10002 // token 验证超时
	CreateTokenError  = 10003 // token创建失败
	TokenNull         = 10004 // 无token数据

	Valerror = 10010 // 验证码错误
	Valright = 10011 // 验证码错误

	UserExits            = 20001 // 用户已存在
	UserNoExits          = 20002 // 用户不存在
	UserCreateError      = 20003 //用户创建失败
	UserPasswordError    = 20004 // 密码错误
	OldPasswordError     = 20005 // 旧密码错误
	EditOldPasswordError = 20006 // 修改密码错误
)
