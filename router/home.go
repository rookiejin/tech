package router

import (
	"gopkg.in/macaron.v1"
	"tech/model"
	"tech/modules/page"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func Home(ctx *macaron.Context)  {
	var serieses []model.Series
	_ = page.Page(ctx , model.SERIES , &bson.M{} , &serieses)
	str , _ := json.Marshal(serieses)
	ctx.Data["series"] = string(str)
	ctx.Data["Name"] = "zhangsan"
	ctx.HTML(200,"home/index")
}
