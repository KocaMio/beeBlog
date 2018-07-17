package controllers

type ErrorController struct {
	BaseController
}

func (e *ErrorController) Error404() {
	e.ResponseJson(JsonResponse {
		Status: 404,
		Msg: "page not found",
	})
}