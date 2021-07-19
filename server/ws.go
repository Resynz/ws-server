/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:36
 */
package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"strconv"
	"time"
	"ws-server/common"
	"ws-server/config"
	"ws-server/lib/logger"
	"ws-server/tools"
)

const (
	XClientID = "X-Client-ID"
	XUserID   = "X-User-ID"
	XPlatform = "X-Platform"
)

var (
	ws = websocket.Server{
		Handshake: HandleShake,
		Handler:   Handler,
	}
)

func putToClientMap(c *config.Client) {
	var cons []*config.Client
	if !config.ClientMap.Exists(c.UserId) {
		cons = make([]*config.Client, 1)
		cons[0] = c
	} else {
		cons = config.ClientMap.Read(c.UserId)
		cons = append(cons, c)
	}
	config.ClientMap.Write(c.UserId, cons)
}

func removeFromClientMap(client *config.Client) {
	client.Conn.Close()
	clients := config.ClientMap.Read(client.UserId)
	var index int
	for i, v := range clients {
		if v == client {
			index = i
		}
	}
	clients = append(clients[:index], clients[index+1:]...)
	if len(clients) > 0 {
		config.ClientMap.Write(client.UserId, clients)
		return
	}
	config.ClientMap.Delete(client.UserId)
}

func Handler(ws *websocket.Conn) {
	pt, _ := strconv.Atoi(ws.Request().Header.Get(XPlatform))
	client := &config.Client{
		Conn:       ws,
		UserId:     ws.Request().Header.Get(XUserID),
		ClientId:   fmt.Sprintf("W_%d_%d", pt, time.Now().UnixNano()),
		Platform:   config.PlatformType(pt),
		CreateTime: time.Now().Unix(),
	}
	putToClientMap(client)
	defer removeFromClientMap(client)
	for {
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Logger.Errorf("用户[%s]数据接收失败！error:%s", client.UserId, err)
			break
		}
	}
}

func HandleShake(conf *websocket.Config, r *http.Request) error {
	userId, platform, err := tools.CheckAuth(r.URL.RawQuery)
	if err != nil {
		return err
	}
	r.Header.Set(XUserID, userId)
	r.Header.Set(XPlatform, fmt.Sprintf("%d", platform))
	return nil
}

func WebSocketHandleFunc(ctx *common.Context) {
	ws.ServeHTTP(ctx.Writer, ctx.Request)
}
