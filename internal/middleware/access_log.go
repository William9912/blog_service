package middleware

import (
	"blog-service/global"
	"blog-service/pkg/logger"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

//一个gin.ResponseWriter的替代品（实现了这个接口）
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//把p的内容写到buffer里
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	//然后再用gin.ResponseWriter里的
	return w.ResponseWriter.Write(p)
}

//一个日志中间件
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		//这个是已经执行完请求了
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}

// //panic中间件
// func Recovery() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
// 				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
// 				c.Abort()
// 			}
// 		}()
// 		c.Next()
// 	}
// }
