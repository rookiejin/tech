package model

import "gopkg.in/mgo.v2/bson"

type Series struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"title"`
	Thumb string `bson:"thumb" json:"thumb"`
	TotalTime int `bson:"total_time" json:"total_time"`
}

