package controllers

import (
	"strconv"
	"time"

	"beeBlog/models"
	"github.com/astaxie/beego/orm"
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
		jsonResponse := JsonResponse {
			Status: 4000,
			Msg: Response["invalidParams"],
		}

		this.ResponseJson(jsonResponse)

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
		jsonResponse := JsonResponse {
			Status: 5000,
			Msg: Response["addDataFaild"],
		}

		this.ResponseJson(jsonResponse)
		return
	} 
	
	data := struct { 
		Id int
	}{
		model.Id,
	}

	jsonResponse := JsonResponse {
		Status: 2000,
		Msg: Response["addDataSuccess"],
		Data: data,
	}

	this.ResponseJson(jsonResponse)
}

func (this *ArticleController) Delete() {
	id, getIdIntegerError := this.GetInt("id")

	// Check Params
	if getIdIntegerError != nil {
		jsonResponse := JsonResponse {
			Status: 4000,
			Msg: Response["invalidParams"],
		}

		this.ResponseJson(jsonResponse)
		return
	}

	// Delete Article
	o := orm.NewOrm()

	model := models.Article{}
	model.Id = id

	if _, error := o.Delete(&model); error != nil {
		jsonResponse := JsonResponse {
			Status: 5000,
			Msg: Response["deleteFaild"],
		}

		this.ResponseJson(jsonResponse)
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
		
		jsonResponse := JsonResponse {
			Status: 4000,
			Msg: Response["invalidParams"],
		}

		this.ResponseJson(jsonResponse)
		return
	}

	// Update by ORM
	o := orm.NewOrm()

	model := models.Article{}
	model.Id = id

	// Check is article exist before update
	if error := o.Read(&model); error != nil {
		jsonResponse := JsonResponse {
			Status: 3000,
			Msg: Response["dataNotFound"],
		}

		this.ResponseJson(jsonResponse)
		return
	} 

	model.Title = title
	model.Content = content
	model.EditTime = time.Now().Format("2006-01-02 15:04:05")

	_, error := o.Update(&model, "Title", "Content", "EditTime")
	
	if error != nil {
		jsonResponse := JsonResponse {
			Status: 5000,
			Msg: Response["modifyFaild"],
		}

		this.ResponseJson(jsonResponse)
		return
	} 
	
	jsonResponse := JsonResponse {
		Status: 2000,
		Msg: Response["modifySuccess"],
	}

	this.ResponseJson(jsonResponse)
}

func (this *ArticleController) GetItem() {
	id := this.GetString("id")

	// Check Post Params
	if "" == id {
		jsonResponse := JsonResponse {
			Status: 2000,
			Msg: Response["invalidParams"],
		}

		this.ResponseJson(jsonResponse)
		return
	}

	idInt, _ := strconv.Atoi(id)

	// Get Item from ORM
	o := orm.NewOrm()
	
	model := models.Article{}
	model.Id = idInt

	if error := o.Read(&model); error != nil {
		jsonResponse := JsonResponse {
			Status: 5000,
			Msg: Response["getDataFaild"],
		}

		this.ResponseJson(jsonResponse)
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
		jsonResponse := JsonResponse {
			Status: 2000,
			Msg: Response["invalidParams"],
		}

		this.ResponseJson(jsonResponse)
		return
	}

	// Response by JSON
	this.Data["json"] = &model
	this.ServeJSON()
}