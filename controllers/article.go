package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"beeBlog/models"
	"strconv"
	"time"
)

type ArticleController struct {
	BaseController
}

type employee struct {
	Id       int
	Name     string
	Salary   int
	Phone    int
	Nickname string
}

func (this *ArticleController) SubPrepare() {

}

func (this *ArticleController) Add() {
	title := this.GetString("title")
	content := this.GetString("content")

	// Check Params
	if "" == title || "" == content {
		fmt.Println("Invalid Params!")
		return
	}

	// Insert by ORM
	o := orm.NewOrm()

	model := models.Article{}
	model.Title = title
	model.Content = content
	model.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	model.EditTime = time.Now().Format("2006-01-02 15:04:05")

	if _, error := o.Insert(&model); error != nil {
		fmt.Println("Insert Faild", error)
		return
	} 

	this.Data["json"] = struct {
		Id int
	}{
		model.Id,
	}

	this.ServeJSON()
}

func (this *ArticleController) Delete() {
	id, getIdIntegerError := this.GetInt("id")

	// Check Params
	if getIdIntegerError != nil {
		fmt.Println("Invalid Params!")
		return
	}

	// Delete Article
	o := orm.NewOrm()

	model := models.Article{}
	model.Id = id

	if _, error := o.Delete(&model); error != nil {
		fmt.Println("Delete Faild")
		return
	}
}

func (this *ArticleController) Update() {
	id, getIdIntegerError := this.GetInt("id")
	title := this.GetString("title")
	content := this.GetString("content")

	// Check Params
	if getIdIntegerError != nil || 
		title == "" || 
		content == "" {
		
		fmt.Println("Invalid Valid!")
	}

	// Update by ORM
	o := orm.NewOrm()

	model := models.Article{}
	model.Id = id

	// Check is article exist before update
	if error := o.Read(&model); error != nil {
		fmt.Println("Article not found")
		return
	} 
// TODO: Fix the problem about createTime cannot restore back from database without any edit
fmt.Println(model.CreateTime)
return
	model.Title = title
	model.Content = content
	model.EditTime = time.Now().Format("2006-01-02 15:04:05")

	number, error := o.Update(&model)
	
	if error != nil {
		fmt.Println(error)
		return
	} 
	
	fmt.Println("Update numbers of data", number)
}

func (this *ArticleController) GetItem() {
	id := this.GetString("id")

	// Check Post Params
	if "" == id {
		fmt.Println("Invalid Params!")
		return
	}

	idInt, _ := strconv.Atoi(id)

	// Get Item from ORM
	o := orm.NewOrm()
	
	model := models.Article{}
	model.Id = idInt

	if error := o.Read(&model); error != nil {
		fmt.Println("Can not get article")
		return
	}

	// Response by JSON
	this.Data["json"] = &model
	this.ServeJSON()
}

func (this *ArticleController) GetList() {

	// Get List
	o := orm.NewOrm()

	model := []models.Article{}

	if _, error := o.QueryTable("article").All(&model); error != nil {
		fmt.Println("SQL Query Error")
		return
	}

	// Response by JSON
	this.Data["json"] = &model
	this.ServeJSON()
}