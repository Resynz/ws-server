/**
 * @Author: Resynz
 * @Date: 2021/7/19 16:51
 */
package api

import (
	"golang.org/x/net/websocket"
	"ws-server/config"
)

func SendMsgsToUsers(users []string, msgs []string) {
	for _, m := range msgs {
		for _, u := range users {
			if config.ClientMap.Exists(u) {
				clients := config.ClientMap.Read(u)
				for _, client := range clients {
					go websocket.Message.Send(client.Conn, m)
				}
			}
		}
	}
}
