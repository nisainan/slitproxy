package user

import (
	"github.com/gin-gonic/gin"
	v1 "slitproxy/user/app/controller"
)

func APIUser(parentRoute gin.IRouter) {
	user := parentRoute.Group("user")
	{
		user.POST("/login", v1.Login)
		user.POST("/register", v1.Register)
	}
}
