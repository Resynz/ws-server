/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:24
 */
package common

import (
	"net/http"
	"ws-server/code"
	"ws-server/config"
)

func HandleResponse(ctx *Context, c code.ResponseCode, d interface{}, msg ...string) {
	m := code.GetCodeMsg(c)
	if len(msg) > 0 {
		m = msg[0]
	}
	if c != code.SuccessCode {
		ctx.Logger.Errorln(m)
	}
	data := map[string]interface{}{
		"code":       c,
		"message":    m,
		"request_id": ctx.GetString(config.Conf.LogConfig.RequestIdKey),
	}
	if d != nil {
		data["data"] = d
	}
	ctx.JSON(http.StatusOK, data)
}
