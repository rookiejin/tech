package seed

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"tech/model"
	"time"
)

func Series(n int) {
	db := model.GetDB()
	for i := 0; i < n; i++ {
		tmp := new(model.Series)
		tmp.Id = bson.NewObjectId()
		tmp.Title = RandomTitle()
		tmp.Description = RandomDescription()
		tmp.Thumb = GetThumb()
		tmp.TotalTime = 0
		tmp.ClickCount = 0
		tmp.Price = 0
		tmp.CreateTime = time.Now()
		tmp.Banner = "http://odcijms8n.bkt.clouddn.com/58c7b9e20001679d16000960.jpg"
		err := db.C("series").Insert(tmp)
		if err != nil {
			log.Fatal("插入失败", err)
		}
	}
}

func Videos(n int) {
	db := model.GetDB()
	var serieses []*model.Series
	db.C("series").Find(bson.M{}).All(&serieses)
	for _, serie := range serieses {
		for i := 0; i < n; i++ {
			tmp := new(model.Videos)
			tmp.Id = bson.NewObjectId()
			tmp.CreateTime = time.Now()
			tmp.Description = RandomDescription()
			tmp.Title = RandomTitle()
			tmp.TotalTime = 20
			tmp.Pid = serie.Id
			tmp.Vid = "10-new-featura-in-laravel-5-5-1507692340366988842.mp4"
			err := db.C("videos").Insert(tmp)
			if err != nil {
				log.Fatal("插入失败", err)
			}
		}
	}
}
