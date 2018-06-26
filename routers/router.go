package routers

import (
	"beeBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.ArticleController{})
	beego.AutoRouter(&controllers.BaseController{})
}
