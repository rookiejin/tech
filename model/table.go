package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// collection name
	SERIES = "series"
	VIDEOS = "videos"
	USERS = "users"
)

type Series struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Thumb       string        `bson:"thumb" json:"thumb"`
	TotalTime   int           `bson:"total_time" json:"total_time"`
	CreateTime  time.Time     `bson:"create_time" json:"create_time"`
}

type Videos struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	TotalTime   int           `bson:"total_time" json:"total_time"`
	CreateTime  time.Time     `bson:"create_time" json:"create_time"`
	Vid         string        `bson:"vid" json:"vid"`
	Pid         bson.ObjectId `bson:"pid" json:"pid"`
}

type Users struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
	Nickname string `bson:"nickname" json:"nickname"`
	HeadImageUrl string `bson:"head_image_url" json:"head_image_url"`
	Github string `bson:"github" json:"github"`
	ProvinceId bson.ObjectId `bson:"province_id" json:"province_id"`
	CityId bson.ObjectId `bson:"city_id" json:"city_id"`
	PersonalSite string `bson:"personal_site" json:"personal_site"`
	Sign string `bson:"sign" json:"sign"`
	InviteCode string `bson:"invite_code" json:"invite_code"`
	Inviter bson.ObjectId `bson:"inviter" json:"inviter"`
	AccessToken string `bson:"access_token" json:"access_token"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
}
