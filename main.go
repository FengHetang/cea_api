/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/17 16:00
 */

package main

import (
	"cea_api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouters(r)

	err := r.Run()
	if err != nil {
		return
	}
}
