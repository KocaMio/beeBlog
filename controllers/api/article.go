package api

import (
    "strconv"
    "time"

    "beeBlog/models"
    . "beeBlog/controllers"
    "github.com/astaxie/beego/orm"
)

type ArticleController struct {
    BaseController
}

func (this *ArticleController) SubPrepare() {

}

func (this *ArticleController) Add() {
    title := this.GetString("title")
    content := this.GetString("content")

    // Check Params
    if "" == title || "" == content {
        this.ResponseJson(JsonResponse {
            Status: 4000,
            Msg: ResponseText["invalidParams"],
        })

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
        this.ResponseJson(JsonResponse {
            Status: 5000,
            Msg: ResponseText["addDataFaild"],
        })

        return
    } 
    
    data := struct { 
        Id int
    }{
        model.Id,
    } 

    this.ResponseJson(JsonResponse {
        Status: 2000,
        Msg: ResponseText["addDataSuccess"],
        Data: data,
    })
}

func (this *ArticleController) Delete() {
    id, getIdIntegerError := this.GetInt("id")

    // Check Params
    if getIdIntegerError != nil {
        this.ResponseJson(JsonResponse {
            Status: 4000,
            Msg: ResponseText["invalidParams"],
        })

        return
    }

    // Delete Article
    o := orm.NewOrm()

    model := models.Article{}
    model.Id = id

    if _, error := o.Delete(&model); error != nil {
        this.ResponseJson(JsonResponse {
            Status: 5000,
            Msg: ResponseText["deleteFaild"],
        })

        return
    }

    this.ResponseJson(JsonResponse {
        Status: 2000,
        Msg: ResponseText["deleteSuccess"],
    })
}

func (this *ArticleController) Update() {
    id, getIdIntegerError := this.GetInt("id")
    title := this.GetString("title")
    content := this.GetString("content")

    // Check Params
    if getIdIntegerError != nil || 
        title == "" || 
        content == "" {

        this.ResponseJson(JsonResponse {
            Status: 4000,
            Msg: ResponseText["invalidParams"],
        })

        return
    }

    // Update by ORM
    o := orm.NewOrm()

    model := models.Article{}
    model.Id = id

    // Check is article exist before update
    if error := o.Read(&model); error != nil {
        this.ResponseJson(JsonResponse {
            Status: 3000,
            Msg: ResponseText["dataNotFound"],
        })

        return
    } 

    model.Title = title
    model.Content = content
    model.EditTime = time.Now().Format("2006-01-02 15:04:05")

    _, error := o.Update(&model, "Title", "Content", "EditTime")
    
    if error != nil {
        this.ResponseJson(JsonResponse {
            Status: 5000,
            Msg: ResponseText["modifyFaild"],
        })

        return
    } 

    this.ResponseJson(JsonResponse {
        Status: 2000,
        Msg: ResponseText["modifySuccess"],
    })
}

func (this *ArticleController) GetItem() {
    id := this.GetString("id")

    // Check Post Params
    if "" == id {
        this.ResponseJson(JsonResponse {
            Status: 4000,
            Msg: ResponseText["invalidParams"],
        })

        return
    }

    idInt, _ := strconv.Atoi(id)

    // Get Item from ORM
    o := orm.NewOrm()
    
    model := models.Article{}
    model.Id = idInt

    if error := o.Read(&model); error != nil {
        this.ResponseJson(JsonResponse {
            Status: 5000,
            Msg: ResponseText["getDataFaild"],
        })

        return
    }

    this.ResponseJson(JsonResponse {
        Status: 2000,
        Msg: ResponseText["getDataSuccess"],
        Data: &model,
    })
}

func (this *ArticleController) GetList() {

    // Get List
    o := orm.NewOrm()

    model := []models.Article{}

    if _, error := o.QueryTable("article").All(&model); error != nil {
        this.ResponseJson(JsonResponse {
            Status: 4000,
            Msg: ResponseText["invalidParams"],
        })
        
        return
    }

    // Response by JSON
    this.ResponseJson(JsonResponse {
        Status: 2000,
        Msg: ResponseText["getDataSuccess"],
        Data: &model,
    })
}