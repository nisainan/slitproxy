package ginpkg

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} { return make([]byte, 0, 32*1024) }}

func NewGin() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	return r
}

func Handle(ctx *gin.Context) {
	httpsProxyURI, err := url.Parse(ctx.GetString("proxy"))
	if err != nil {
		log.Printf("failed to parse https proxy uri : %s\n", err)
		return
	}
	dialer, err := proxy.FromURL(httpsProxyURI, proxy.Direct)
	if err != nil {
		log.Printf("failed to proxy.FromURL : %s\n", err)
		return
	}
	// TODO default port is 80
	connProxy, err := dialer.Dial("tcp4", fmt.Sprintf("%s:%d", ctx.Request.Host, 80))
	if err != nil {
		log.Println("httpsDialer.Dial error:", err)
		return
	}
	forwardResponse(connProxy, ctx.Request, ctx.Writer)
}

func forwardResponse(conn net.Conn, req *http.Request, res http.ResponseWriter) {
	err := req.Write(conn)
	if err != nil {
		log.Printf("failed to write http request : %s\n", err)
		return
	}
	response, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		log.Printf("failed to read http response : %s\n", err)
		return
	}
	req.Body.Close()
	if response != nil {
		defer response.Body.Close()
	}

	for header, values := range response.Header {
		for _, val := range values {
			res.Header().Add(header, val)
		}
	}
	//util.RemoveHopByHop(res.Header())
	res.WriteHeader(response.StatusCode)
	buf := bufferPool.Get().([]byte)
	buf = buf[0:cap(buf)]
	io.CopyBuffer(res, response.Body, buf)
}
