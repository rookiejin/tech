package model

import "gopkg.in/mgo.v2/bson"

func GetSeriesById( id bson.ObjectId , result *Series) error {
	db := GetDB()
	err := db.C(SERIES).Find(&bson.M{
		"_id" : id ,
	}).One(result)
	return err
}