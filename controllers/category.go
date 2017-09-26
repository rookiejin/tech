package controllers

import (
	"tech/models"
)

type CategoryController struct {
	BaseController
}
/**
	视频分类列表 .
 */
func (c *CategoryController) Get(){
	size , index := Pagition(c)
	cate := new(models.Category)
	_,category := cate.List(index ,size)
	c.Success(category)
}


func Pagition(c *CategoryController) (int ,int) {
	size , err := c.GetInt("size" , 10)
	if err != nil {
		size = 10
	}
	index , err := c.GetInt("index", 0)
	if err != nil {
		index = 0
	}
	return size ,index
}