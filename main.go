package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"slitproxy/ginpkg"
	_ "slitproxy/proxy"
)

var port = flag.Int("port", 80, "Input Port")
var proxy = flag.String("proxy", "", "https proxy url")

func init() {
	flag.Parse()
}

func main() {
	engine := ginpkg.NewGin()
	engine.Use(func(c *gin.Context) {
		c.Set("proxy", *proxy)
	})
	engine.Use(ginpkg.Handle)
	err := engine.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
}
