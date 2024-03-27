// 验证码 控制层
// author xiaoRui

package controller

import (
	"admin-api/api/service"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// @Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
