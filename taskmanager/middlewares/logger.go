package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"
)

func GinLogger(logger *zap.Logger) gin.HandlerFunc { //基于Gin实现的Zap中间件
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", raw),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
		//utils.InitZap()
	}

}

// // GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//检查连接是否断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					//err断言为*net.OpError类型。net.OpError是Go的net包中定义的一个错误类型，用于表示网络操作中的错误。
					//如果断言成功（即err确实是一个*net.OpError类型的错误），ok将为true，并且ne将包含err的*net.OpError值。
					if se, ok := ne.Err.(net.Error); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true //全部条件都通过，即表示错误为网络连接异常
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				//生成一个易于类人阅读的 HTTP 请求文本表示
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					//网络异常
					_ = c.Error(err.(error))
					c.Abort()
					return

				}
				if stack { //根据stack 来决定是否返回堆栈信息
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())))
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
