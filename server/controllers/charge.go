package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 房产相关操作处理
type ChargeControllers struct {
	beego.Controller
}

// 收费项目信息表
func (h *ChargeControllers) GetChangeItemList() {
	// 查询数据
	o := orm.NewOrm()
	rows := new([]models.EsExpensesProject)
	_, err := o.QueryTable(models.TableEsExpensesProject).All(rows)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsExpensesProject, err)
	}

	// 获取响应结果
	response := utils.Response{
		Code: 20000,
		Data: &utils.ResponseTable{
			Total: len(*rows),
			Items: rows,
		},
	}
	h.Data["json"] = response
	h.ServeJSON()
}

// 费用信息表
func (h *ChargeControllers) GetChangeRecordList() {
	// 查询数据
	o := orm.NewOrm()
	rows := new([]models.EsExpenses)
	_, err := o.QueryTable(models.TableEsExpenses).All(rows)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsExpenses, err)
	}

	// 获取响应结果
	response := utils.Response{
		Code: 20000,
		Data: &utils.ResponseTable{
			Total: len(*rows),
			Items: rows,
		},
	}
	h.Data["json"] = response
	h.ServeJSON()
}
