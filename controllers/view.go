package controllers

type ViewController struct {
	BaseController
}

func (this *ViewController) Index() {
	this.TplName = "index.tpl"
}

