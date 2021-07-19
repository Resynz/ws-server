/**
 * @Author: Resynz
 * @Date: 2021/7/19 13:48
 */
package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ws-server/config"
	"ws-server/lib/logger"
)

type Context struct {
	*gin.Context
	Logger *logrus.Entry
}

type HandlerFunc func(ctx *Context)

func AuthDetection(next HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := new(Context)
		context.Context = ctx

		context.Logger = logger.Logger.WithFields(logrus.Fields{
			"request_id": ctx.GetString(config.Conf.LogConfig.RequestIdKey),
		})
		next(context)
	}
}
