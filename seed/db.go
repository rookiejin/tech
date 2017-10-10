package seed

import "tech/model"

func Series()  {
	db := model.GetDB()
	db.C("series").Insert()
}