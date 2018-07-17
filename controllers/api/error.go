package api

import (
	. "beeBlog/controllers"
)

type ErrorController struct {
	BaseController
}

func (e *ErrorController) Error404() {
	e.Ctx.Redirect(302, "/")
}