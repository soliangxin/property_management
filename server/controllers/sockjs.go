package controllers

import "github.com/astaxie/beego"

type SockJsController struct {
	beego.Controller
}

func (s *SockJsController) Get() {
	response := make(map[string]interface{})
	responseData := make(map[string]interface{})

	response["code"] = 20000
	response["data"] = responseData
	s.Data["json"] = response
	s.ServeJSON()
}
