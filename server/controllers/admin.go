package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录
func (l *LoginController) Login() {
	var user User
	data := l.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	response := make(map[string]interface{})
	responseData := make(map[string]interface{})
	responseData["token"] = "admin-token"
	response["code"] = 20000
	response["data"] = responseData
	l.Data["json"] = response
	l.ServeJSON()
}

// 获取用户信息
func (l *LoginController) GetInfo() {
	response := make(map[string]interface{})
	responseData := make(map[string]interface{})
	responseData["roles"] = []string{"admin"}
	responseData["introduction"] = "我是超级管理员"
	responseData["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	responseData["name"] = "超级管理员"

	response["code"] = 20000
	response["data"] = responseData
	l.Data["json"] = response
	l.ServeJSON()
}

// 用户退出
func (l *LoginController) Logout() {
	response := make(map[string]interface{})
	response["code"] = 20000
	response["data"] = "success"
	l.Data["json"] = response
	l.ServeJSON()
}
