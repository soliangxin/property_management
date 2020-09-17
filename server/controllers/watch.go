package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 值班人员管理
type WatchControllers struct {
	beego.Controller
}

// 获取全部小区信息
func (h *WatchControllers) GetWatchList() {
	// 查询数据
	o := orm.NewOrm()
	house := new([]models.EsDuty)
	_, err := o.QueryTable(models.TableEsDuty).All(house)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsDuty, err)
	}

	// 获取响应结果
	response := utils.Response{
		Code: 20000,
		Data: &utils.ResponseTable{
			Total: len(*house),
			Items: house,
		},
	}
	h.Data["json"] = response
	h.ServeJSON()
}
