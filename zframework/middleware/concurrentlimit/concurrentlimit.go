package concurrentlimit

import (
	"errors"
	"net/http"
	"shadowDemo/zframework/utils"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

//并发数map，key为ip，value为当前并发数
var concurrentHolder sync.Map

//黑名单，key为ip，value为到期时间
var blackList sync.Map

type ConcurrentLimitError struct {
	Err error
}

func (e ConcurrentLimitError) Error() string {
	return e.Err.Error()
}

func ConcurrentLimit(limit int, holdDuration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := utils.GetRealIp(c.Request)
		if holdTime, ok := blackList.Load(ip); ok {
			//如果到期时间未到， 则直接返回错误, 否则从黑名单里删除
			if time.Now().Before(holdTime.(time.Time)) {
				// Log.WithField("IP", ip).Warn("ip is in concurrent limit black list")
				err := ConcurrentLimitError{
					Err: errors.New("Request too often"),
				}
				c.AbortWithError(http.StatusForbidden, err)
				return
			} else {
				blackList.Delete(ip)
			}
		}

		if counter, ok := concurrentHolder.Load(ip); ok {
			count := counter.(*int32)
			//如果并发数未达到限制，则放行，否则判定holdMins是否大于0， 如果大于0，则加入黑名单，否则直接返回错误
			if int(*count) < limit {
				atomic.AddInt32(count, 1)
				defer atomic.AddInt32(count, -1)
				c.Next()
			} else {
				Log.WithField("IP", ip).Warnf("ip concurrent count exceed limit %d", limit)
				if holdDuration > 0 {
					blackList.Store(ip, time.Now().Add(holdDuration))
				}
				err := ConcurrentLimitError{
					Err: errors.New("Request too often"),
				}
				c.AbortWithError(http.StatusForbidden, err)
				return
			}
		} else {
			var count = int32(0)
			concurrentHolder.Store(ip, &count)
			c.Next()
		}
	}
}
