package internal

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewRouter(backEnd string) *gin.Engine {
	backEndServer := fmt.Sprintf("http://%s", backEnd)
	targetURL, err := url.Parse(backEndServer)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	defaultRouter := gin.Default()
	defaultRouter.Match([]string{http.MethodGet, http.MethodPost}, "/user-management", func(ctx *gin.Context) {
		ctx.Request.URL.Scheme = targetURL.Scheme
		ctx.Request.URL.Host = targetURL.Host
		ctx.Request.Host = targetURL.Host
		// 将请求转发给后端服务器，怎么加负载均衡？
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})
	return defaultRouter
}
