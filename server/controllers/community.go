// 查询小区信息
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"server/models"
	"server/utils"
)

// 小区相关操作处理
type CommunityControllers struct {
	beego.Controller
}

// 获取全部小区信息
func (c *CommunityControllers) GetCommunityList() {
	// 查询数据
	o := orm.NewOrm()
	community := new([]models.EsCommunity)
	_, err := o.QueryTable(models.TableEsCommunity).All(community)
	if err != nil {
		logs.Error("query %s error, error message: %s", models.TableEsCommunity, err)
	}

	// 获取响应结果
	response := utils.Response{
		Code: 20000,
		Data: &utils.ResponseTable{
			Total: len(*community),
			Items: community,
		},
	}
	c.Data["json"] = response
	c.ServeJSON()
}
