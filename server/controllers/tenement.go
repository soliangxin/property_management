package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 住户信息管理
type TenementControllers struct {
	beego.Controller
}

// 获取全部小区信息
func (t *TenementControllers) GetTenementList() {
	// 查询数据
	o := orm.NewOrm()
	house := new([]models.EsMember)
	_, err := o.QueryTable(models.TableEsMember).All(house)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsMember, err)
	}

	// 获取响应结果
	response := utils.Response{
		Code: 20000,
		Data: &utils.ResponseTable{
			Total: len(*house),
			Items: house,
		},
	}
	t.Data["json"] = response
	t.ServeJSON()
}
