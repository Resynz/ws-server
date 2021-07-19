/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:33
 */
package main

import (
	"log"
	"ws-server/config"
	"ws-server/lib/logger"
	"ws-server/server"
)

func main() {
	log.Println("initializing logger ...")
	if err := logger.InitLogger(config.Conf.LogConfig.Path, config.Conf.LogConfig.Name, config.Conf.LogConfig.Level); err != nil {
		log.Fatalf("init logger error:%s\n", err.Error())
	}
	log.Println("\033[42;30m DONE \033[0m[WS-Server] Start Success!")
	server.StartServer()
}
