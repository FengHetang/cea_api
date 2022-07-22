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

type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}

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
