package controllers

import (
	"beeBlog/library/language"
	"github.com/astaxie/beego"
)

type JsonResponse struct {
	Status int 			`json:"status"`
	Msg string			`json:"msg"`
	Data interface{}	`json:"data"`
}

type SubPreparer interface {
	SubPrepare()
}

type BaseController struct {
	beego.Controller
}

var ResponseText map[string]string

func init() {

	// Init Response from language
	ResponseText = language.New("zhTw").Response
}

func (this *BaseController) Prepare() {
	if app, ok := this.AppController.(SubPreparer); ok {
		app.SubPrepare()
	}
}

func (this *BaseController) ResponseJson(jsonResponse JsonResponse) {
	this.Data["json"] = jsonResponse

	this.ServeJSON()
}
