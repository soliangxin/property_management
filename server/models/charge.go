// 收费管理
// 主要功能模块：收费项目定义、缴费记录管理
package models

import "time"

const (
	TableEsExpensesProject = "es_expenses_project" // 收费项目信息表
	TableEsExpenses        = "es_expenses"         // 费用信息表
)

// 收费项目信息表
type EsExpensesProject struct {
	Id            int       `orm:"column(Id)" json:"id"`                           // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`                 // 所属小区编号，建议CM开头
	Code          string    `orm:"size(50)" json:"code"`                           // 项目编号，建议EP开头
	Name          string    `orm:"size(255)" json:"name"`                          // 项目名称
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"` // 添加时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"update_time"`     // 修改时间
}

// 费用信息表
type EsExpenses struct {
	Id            int       `orm:"column(Id)" json:"id"`                       // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`             // 所属小区编号，建议CM开头
	HouseCode     string    `orm:"size(50)" json:"house_code"`                 // 缴费人，即房产编号，建议HS开头
	ProjectCode   string    `orm:"size(50)" json:"project_code"`               // 缴费项目编号，建议EP开头
	AmountTotal   float64   `orm:"digits(10);decimals(2)" json:"amount_total"` // 应收金额
	AmountPaid    float64   `orm:"digits(10);decimals(2)" json:"amount_paid"`  // 实收金额
	PayTime       time.Time `orm:"auto_now;type(datetime)" json:"pay_time"`    // 缴费时间
	Remark        string    `orm:"size(255)" json:"remark"`                    // 备注
}
