package router

import (
	"tech/modules/middleware"
	"encoding/hex"
)


func IdFilter(ctx *middleware.Context){
	id := ctx.Query("id")
	if id != "" {
		h , err  := hex.DecodeString(id)
		if err!= nil && len(h) != 12 {
			ctx.Data["status_code"] = 404
			ctx.Data["status"] = "fail"
			ctx.Data["message"] = "404 not found"
		}
	}
}
