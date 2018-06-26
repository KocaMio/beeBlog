package controllers

import (
	"fmt"
	"encoding/json"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) SubPrepare() {
	
}

func (this *ArticleController) Read() {
	body := this.Ctx.Input.RequestBody
	
	result := map[string]interface{}{}

	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
