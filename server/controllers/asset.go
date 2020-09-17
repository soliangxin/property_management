package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 资产相关处理
type AssetControllers struct {
	beego.Controller
}

// 获取全部小区信息
func (h *AssetControllers) GetAssetList() {
	// 查询数据
	o := orm.NewOrm()
	house := new([]models.EsDevice)
	_, err := o.QueryTable(models.TableEsDevice).All(house)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsDevice, err)
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
