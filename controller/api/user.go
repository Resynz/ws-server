/**
 * @Author: Resynz
 * @Date: 2021/7/19 16:41
 */
package api

import (
	"ws-server/code"
	"ws-server/common"
	"ws-server/config"
	"ws-server/server"
)

// GetOnlineCount 获取在线人数
func GetOnlineCount(ctx *common.Context) {
	data := map[string]int{
		"count": config.ClientMap.Size(),
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// IsOnline 检测是否在线
func IsOnline(ctx *common.Context) {
	type formValidate struct {
		UserId string `form:"user_id" binding:"required" json:"user_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	result := config.ClientMap.Exists(form.UserId)
	data := map[string]bool{
		"result": result,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// OnlineUsers 获取在线用户列表
func OnlineUsers(ctx *common.Context) {
	list := config.ClientMap.Keys()
	data := map[string][]string{
		"list": list,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// UserInfo 获取在线用户详情
func UserInfo(ctx *common.Context) {
	type formValidate struct {
		UserId string `form:"user_id" binding:"required" json:"user_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}

	type clientObj struct {
		ClientId   string              `json:"client_id"`
		CreateTime int64               `json:"create_time"`
		Platform   config.PlatformType `json:"platform"`
	}

	type infoObj struct {
		UserId  string       `json:"user_id"`
		Clients []*clientObj `json:"clients"`
	}

	info := &infoObj{
		UserId:  form.UserId,
		Clients: nil,
	}

	if config.ClientMap.Exists(form.UserId) {
		cs := config.ClientMap.Read(form.UserId)
		clients := make([]*clientObj, len(cs))
		for i, v := range cs {
			clients[i] = &clientObj{
				ClientId:   v.ClientId,
				CreateTime: v.CreateTime,
				Platform:   v.Platform,
			}
		}
		info.Clients = clients
	}
	data := map[string]interface{}{
		"info": info,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// Offline 下线
func Offline(ctx *common.Context) {
	type formValidate struct {
		UserId   string `form:"user_id" binding:"required" json:"user_id"`
		ClientId string `form:"client_id" binding:"required" json:"client_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	if config.ClientMap.Exists(form.UserId) {
		clients := config.ClientMap.Read(form.UserId)
		for _, v := range clients {
			if v.ClientId == form.ClientId {
				server.RemoveFromClientMap(v)
				break
			}
		}
	}
	data := map[string]bool{
		"result": true,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}
