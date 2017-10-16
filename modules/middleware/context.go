package middleware
//
//import (
//	"fmt"
//	"github.com/qiniu/log"
//	"gopkg.in/macaron.v1"
//)
//
//type Context struct {
//	*macaron.Context
//}
//
//func Contexter() macaron.Handler {
//	return func(c *macaron.Context) {
//		ctx := &Context{
//			c,
//		}
//		c.Map(ctx)
//		defer func(ctx *Context) {
//			env := macaron.Env
//			if err := recover(); err != nil {
//				if env != macaron.PROD {
//					log.Error(err)
//				}
//				if fmt.Sprintf("%v", err) == "not found" {
//					ctx.Data["status_code"] = 404
//					ctx.Data["status"] = "fail"
//					ctx.Data["message"] = "404 NotFound"
//					ctx.Response()
//					return
//				}
//				ctx.Data["status_code"] = 500
//				ctx.Data["status"] = "fail"
//				ctx.Data["message"] = "Server Occur Error"
//				ctx.Response()
//			}
//		}(ctx)
//		ctx.Next()
//		ctx.Response()
//	}
//}
//
//func (ctx *Context) GetStatusCode() int {
//	code, ok := ctx.Data["status_code"]
//	if !ok {
//		return 200
//	}
//	return code.(int)
//}
//
//func (ctx *Context) Response(v ...interface{}) {
//	status_code := ctx.GetStatusCode()
//	status, ok := ctx.Data["status"]
//	if !ok {
//		status = "success"
//	}
//	message, ok := ctx.Data["message"]
//	if !ok {
//		message = ""
//	}
//	data, ok := ctx.Data["data"]
//	if !ok {
//		data = ""
//	}
//	ctx.JSON(status_code, map[string]interface{}{
//		"status":  status,
//		"message": message,
//		"data":    data,
//	})
//}
