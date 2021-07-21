/**
 * @Author: Resynz
 * @Date: 2021/7/19 16:44
 */
package api

import (
	"ws-server/code"
	"ws-server/common"
	"ws-server/config"
)

// SendMsg 发送消息
func SendMsg(ctx *common.Context) {
	type formValidate struct {
		MsgList      []string `form:"msg_list" binding:"required" json:"msg_list"`
		UserIdList   []string `form:"user_id_list" binding:"required" json:"user_id_list"`
		ClientIdList []string `form:"client_id_list" binding:"" json:"client_id_list"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	SendMsgsToUsers(form.UserIdList, form.MsgList, form.ClientIdList)
	data := map[string]bool{
		"result": true,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// Broadcast 广播
func Broadcast(ctx *common.Context) {
	type formValidate struct {
		MsgList []string `form:"msg_list" binding:"required" json:"msg_list"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	users := config.ClientMap.Keys()
	if len(users) > 0 {
		SendMsgsToUsers(users, form.MsgList, nil)
	}
	data := map[string]bool{
		"result": true,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}
