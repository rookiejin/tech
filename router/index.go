package router

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
	"tech/model"
	"tech/modules/page"
)

// 首页的所有视频
func All(ctx *macaron.Context) {
	var serieses []model.Series
	var err error
	err = page.Page(ctx, model.SERIES, &bson.M{}, &serieses)
	if err != nil {
		ctx.Data["message"] = fmt.Sprintf("%v", err)
		log.Error(err)
	} else {
		ctx.Data["data"] = serieses
	}
}
