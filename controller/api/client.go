/**
 * @Author: Resynz
 * @Date: 2021/7/19 16:51
 */
package api

import (
	"golang.org/x/net/websocket"
	"ws-server/config"
	"ws-server/tools"
)

func SendMsgsToUsers(users []string, msgs, clientIds []string) {
	for _, m := range msgs {
		for _, u := range users {
			if config.ClientMap.Exists(u) {
				clients := config.ClientMap.Read(u)
				for _, client := range clients {
					if clientIds == nil || len(clientIds) == 0 {
						go websocket.Message.Send(client.Conn, m)
					}
					if tools.CheckIsInStringArray(clientIds, client.ClientId) {
						go websocket.Message.Send(client.Conn, m)
					}
				}
			}
		}
	}
}
