package controller

import (
	"slitproxy/user/app/base/controller"
	"slitproxy/user/app/model/mparam"
	"slitproxy/user/app/service"

	"github.com/gin-gonic/gin"
	"slitproxy/user/pkg/response"
)

func Register(c *gin.Context) {
	param := &mparam.Register{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.Register(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

func Login(c *gin.Context) {
	param := &mparam.Login{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.Login(c, param)
	response.UtilResponseReturnJson(c, code, data)
}
