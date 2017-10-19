package router

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2/bson"
	"tech/model"
	"tech/modules/middleware"
	"tech/modules/page"
	"fmt"
)

func Home(ctx *macaron.Context) {
	var serieses []model.Series
	_ = page.Page(ctx, model.SERIES, &bson.M{}, &serieses)
	str, _ := json.Marshal(serieses)
	ctx.Data["series"] = string(str)
	ctx.Data["Title"] = "首页"
	ctx.HTML(200, "home/index")
}

func Detail(ctx *macaron.Context) {
	defer middleware.ServerIdValidFilter(ctx)
	id := ctx.Query("id")
	var (
		video  []model.Videos
		err    error
		series model.Series
		Id     bson.ObjectId
	)
	Id = bson.ObjectIdHex(id)
	err = model.GetSeriesById(Id, &series)
	if err != nil {
		ctx.HTML(404, "error/404")
	}
	err = page.Page(ctx, model.VIDEOS, &bson.M{"pid": Id}, &video)
	if err != nil {
		ctx.HTML(404, "error/404")
		return
	}
	ctx.Data["data"] = map[string]interface{}{
		"series": series,
		"videos": video,
	}
	ctx.Data["Title"] = series.Title
	ctx.HTML(200, "home/detail")
}

// 视频页面
func Video(ctx *macaron.Context) {
	//defer middleware.ServerIdValidFilter(ctx)
	id := ctx.Query("id")
	var (
		video *model.Videos = &model.Videos{}
		err   error
		Id    bson.ObjectId
	)
	Id = bson.ObjectIdHex(id)
	err = model.GetVideoById(Id, video)
	// 阅读 + 1
	model.AddVideoClickCount(Id)
	fmt.Println("xxx2")

	if err != nil {
		ctx.HTML(404, "error/404")
		return
	}

	ctx.Data["data"] = video
	ctx.HTML(200, "home/video")
}
