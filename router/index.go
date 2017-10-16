package router

import (
	"tech/model"
	"tech/modules/page"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
	"fmt"
	"gopkg.in/macaron.v1"
)
// 首页的所有视频
func All(ctx *macaron.Context){
	var serieses []model.Series
	var err error
	err = page.Page(ctx , model.SERIES , &bson.M{} , &serieses)
	if err != nil {
		ctx.Data["message"] = fmt.Sprintf("%v",err)
		log.Error(err)
	}else{
		ctx.Data["data"] = serieses
	}
}
// 详情
func Detail(ctx *macaron.Context) {
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
		ctx.Data["status_code"] = 404
		ctx.Data["status"] = "fail"
		ctx.Data["message"] = fmt.Sprintf("%v",err)
		return
	}
	err = page.Page(ctx ,model.VIDEOS , &bson.M{"pid" :Id},&video)
	if err != nil {
		ctx.Data["status_code"] = 404
		ctx.Data["status"] = "fail"
		ctx.Data["message"] = fmt.Sprintf("%v",err)
		return
	}
	ctx.Data ["data"] = map[string]interface{}{
		"series" : series ,
		"videos" : video ,
	}
}