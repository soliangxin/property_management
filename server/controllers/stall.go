package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 房产相关操作处理
type StallControllers struct {
	beego.Controller
}

// 获取全部小区信息
func (h *StallControllers) GetStallList() {
	// 查询数据
	o := orm.NewOrm()
	house := new([]models.EsParkingSpace)
	_, err := o.QueryTable(models.TableEsParkingSpace).All(house)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsParkingSpace, err)
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
