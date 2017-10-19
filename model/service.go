package model

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func GetSeriesById(id bson.ObjectId, result *Series) error {
	db := GetDB()
	err := db.C(SERIES).Find(&bson.M{
		"_id": id,
	}).One(result)
	return err
}

func GetVideoById(id bson.ObjectId, result *Videos) error {
	db := GetDB()
	err := db.C(VIDEOS).Find(&bson.M{
		"_id": id,
	}).One(result)
	return err
}

func AddVideoClickCount(id bson.ObjectId)  {
	fmt.Println("xxx")
	db := GetDB()
	db.C(VIDEOS).UpdateId(id , bson.M{
		"$inc" : bson.M{
			"click_count" : +1 ,
		},
	})
}