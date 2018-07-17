package routers

import (
    "beeBlog/controllers/api"
    "beeBlog/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.AutoRouter(&api.ArticleController{})
    beego.Router("/", &controllers.ViewController{}, "*:Index")
}
