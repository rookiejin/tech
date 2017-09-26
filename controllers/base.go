package controllers

import (
	"github.com/astaxie/beego"
	"tech/utils"
)

type BaseController struct{
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
}

func (c *BaseController) Success ( data interface{} , args ...interface{}) {
	json := utils.Json{
		"data" : data ,
	}
	l := len(args)
	if l > 0 {
		json ["message"] = args [0]
	}
	json ["status"] = "ok"
	c.Data["json"] = json
	c.ServeJSON()
}
/**
 * 返回错误信息
 */
func (c *BaseController) Error ( message string , args ...interface{}) {
	json := utils.Json{
		"message" : message ,
	}
	l := len(args)
	if l > 0 {
		json ["data"] = args [1]
	}
	json ["status"] = "fail"
	c.Data["json"] = json
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.ServeJSON()
}