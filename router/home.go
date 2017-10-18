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
	ctx.Data["Title"] = "首页"
	ctx.HTML(200,"home/index")
}

func Detail(ctx *macaron.Context)  {
	id := ctx.Query("id")
	var (
		video []model.Videos
		err error
		series model.Series
		Id bson.ObjectId
	)
	Id = bson.ObjectIdHex(id)
	err = model.GetSeriesById(Id , &series)
	if err != nil {
		ctx.HTML(404,"error/404")
	}
	err = page.Page(ctx ,model.VIDEOS , &bson.M{"pid" :Id},&video)
	if err != nil {
		ctx.HTML(404,"error/404")
	}
	ctx.Data ["data"] = map[string]interface{}{
		"series" : series ,
		"videos" : video ,
	}
	ctx.Data["Title"] = series.Title
	ctx.HTML(200,"home/detail")
}