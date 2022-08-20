package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// Key：获取对应的限流器的键值对名称。
// GetBucket：获取令牌桶。
// AddBuckets：新增多个令牌桶。
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

//一个字符串和ratelimit.Bucket的字典
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// Key：自定义键值对名称。
// FillInterval：间隔多久时间放 N 个令牌。
// Capacity：令牌桶的容量。
// Quantum：每次到达间隔时间后所放的具体令牌数量。
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
