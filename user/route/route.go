package route

import (
	"slitproxy/user/app/base/controller"
	"slitproxy/user/pconst"
	"slitproxy/user/pkg/confer"
	"slitproxy/user/route/user"

	"github.com/gin-gonic/gin"
)

func Home(parentRoute *gin.Engine) {
	parentRoute.GET("", controller.Welcome)
}

func Api(engine *gin.Engine) {
	prefix := confer.ConfigAppGetString("UrlPrefix", "")
	RouteV1 := engine.Group(prefix + pconst.APIAPIV1URL)
	{
		user.APIUser(RouteV1)
	}
}

func NotFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(404, "404 Not Found")
	})
}
