package models

import (
    "github.com/astaxie/beego/orm"
)

type Article struct {
    Id			int 	`orm:"column(id)"`
    Title		string	`orm:"column(title)"`
    Content		string 	`orm:"column(content)"`
    CreateTime	string	`orm:"column(createTime);auto_now_add"`
    EditTime	string	`orm:"column(editTime);auto_now"`
    Tag			string	`orm:"column(tag);type(json)"`
}

func init() {
    orm.RegisterModel(new(Article))
}