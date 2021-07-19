/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:26
 */
package controller

import (
	"ws-server/code"
	"ws-server/common"
)

func Ping(ctx *common.Context) {
	ctx.Logger.Infoln("八嘎", "小八嘎")
	common.HandleResponse(ctx, code.SuccessCode, nil)
}
