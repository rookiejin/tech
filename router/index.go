package router

import (
	"gopkg.in/macaron.v1"
	"tech/model"
	"tech/modules/page"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
	"fmt"
	"encoding/json"
)


// 首页的所有视频
func All(ctx *macaron.Context){
	var serieses []model.Series
	var err error
	err = page.Page(ctx , model.SERIES , &bson.M{} , &serieses)
	if err != nil {
		log.Error(err)
	}
	ctx.JSON(200 , serieses)
}
// 详情
func Detail(ctx *macaron.Context) {
	id := ctx.Query("id")
	fmt.Printf(" origin %s",id)
	var video []model.Videos
	var err error
	err = page.Page(ctx ,model.VIDEOS , &bson.M{"pid" :bson.ObjectIdHex(id)},&video)
	if err != nil {
		e :=  make(map[string]string)
		e["message"] = "404 not found"
		ctx.JSON(404, e)
	}
	ctx.JSON(200 ,video)
}