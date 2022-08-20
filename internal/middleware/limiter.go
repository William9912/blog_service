package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"blog-service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		//这里是只限制能GetBucket拿到的uri 如果拿不到 则不限制
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			//减少一个count
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		//这 如果拿不到 则不限制 走这里
		c.Next()
	}
}
