/**
 * @Author: Resynz
 * @Date: 2021/5/26 18:09
 */
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
	"ws-server/config"
	"ws-server/lib/logger"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 生成requestId
		requestId := fmt.Sprintf("RQ%d", time.Now().UnixNano())
		c.Set(config.Conf.LogConfig.RequestIdKey, requestId)

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime).Nanoseconds() / 1e6

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Logger.WithFields(logrus.Fields{
			"request_id":   requestId,
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()

	}
}
