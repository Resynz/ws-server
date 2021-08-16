/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:29
 */
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ws-server/common"
	"ws-server/config"
	"ws-server/controller"
	"ws-server/controller/api"
	"ws-server/middleware"
)

func StartServer() {
	gin.SetMode(config.Conf.Mode)
	app := gin.New()
	app.MaxMultipartMemory = 8 << 20 // 8mb
	app.Use(gin.Recovery())
	app.Use(middleware.Logger())

	app.GET("/ping", common.AuthDetection(controller.Ping))

	app.GET("/ws", common.AuthDetection(WebSocketHandleFunc))

	apiGroup := app.Group("/api")
	apiGroup.GET("/online-count", common.AuthDetection(api.GetOnlineCount))
	apiGroup.POST("/send-msg", common.AuthDetection(api.SendMsg))
	apiGroup.POST("/broadcast", common.AuthDetection(api.Broadcast))
	apiGroup.GET("/is-online", common.AuthDetection(api.IsOnline))
	apiGroup.GET("/online-users", common.AuthDetection(api.OnlineUsers))
	apiGroup.GET("/info", common.AuthDetection(api.UserInfo))
	apiGroup.GET("/offline", common.AuthDetection(api.Offline))

	if err := app.Run(fmt.Sprintf(":%d", config.Conf.AppPort)); err != nil {
		log.Fatalf("start server failed! error:%s\n", err.Error())
	}
}
