package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Category struct {
	Id int `json:"id" orm:"pk;auto;"`
	Name string `json:"name" orm:"size(50)"`
	Pid int `json:"pid"`
	Thumb string `json:"thumb" orm:"size(255)"`
	Description string `json:"description" orm:"type(text)"`
	Total_time int `json:"total_time"`
	Nums int `json:"nums"`
	Created_at time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}


func (c *Category) List(index int ,size int) (int64,[]*Category){
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	qs.OrderBy("-id")
	var res []*Category
	num ,err := o.QueryTable(c).OrderBy("-id").Limit(size,index).All( &res )
	if err != nil {
		fmt.Println("查询失败")
	}
	return num , res
}
