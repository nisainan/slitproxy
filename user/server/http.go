package server

import (
	"fmt"
	"strconv"

	"slitproxy/user/pkg/confer"
	"slitproxy/user/pkg/gin"
	"slitproxy/user/pkg/middle"
	"slitproxy/user/route"
)

func RunHTTP() {
	engine := gin.NewGin()
	//engine.Use(middle.RequestID())
	// 仅针对开发环境生效的组件
	//if confer.ConfigEnvIsDev() {
	// 跨域中间件
	engine.Use(middle.CorsV2())
	// swagger
	//}
	route.Home(engine)
	route.Api(engine)
	route.NotFound(engine)
	httpPort := confer.ConfigAppGetInt("port", 80)
	portStr := ":" + strconv.Itoa(httpPort)
	fmt.Println("start", httpPort)
	gin.ListenHTTP(portStr, engine, 10)
}
