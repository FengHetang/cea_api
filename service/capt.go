/**
 * @Author: 戈亓
 * @Description:
 * @File:  capt
 * @Version: 1.0.0
 * @Date: 2022/7/22 15:03
 */

package service

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (id string, b64s string) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()

	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", ""
	}
	return id, b64s
}
func VerityCaptcha(id string, ret_captcha string) bool {
	return store.Verify(id, ret_captcha, true)
}
