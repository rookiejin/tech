package page

import (
	"tech/model"
	"tech/modules/setting"
	"gopkg.in/macaron.v1"
)

func Page(ctx *macaron.Context, table string, query interface{} , getter interface{}) error {
	db := model.GetDB()
	pageNow := ctx.QueryInt("page")
	if pageNow <= 0 {
		pageNow = 1
	}
	err := db.C(table).Find(&query).Limit(setting.PageSize).Skip((pageNow - 1) * setting.PageSize).All(getter)
	return err
}
