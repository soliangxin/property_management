package models

import "time"

const (
	TableEsDevice = "es_device"
)

// 资产设备信息表
type EsDevice struct {
	Id            int       `orm:"column(Id)" json:"id"`                    // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`          // 所属小区编号，建议CM开头
	Code          string    `orm:"size(50)" json:"code"`                    // 设备编号
	Name          string    `orm:"size(50)" json:"name"`                    // 设备名称
	Brand         string    `orm:"size(50)" json:"brand"`                   // 品牌
	Price         float64   `orm:"digits(10);decimals(2)" json:"price"`     // 购买价格（单价）
	Quantity      int       `orm:"size(10)" json:"quantity"`                // 购买数量
	DurableYears  int       `orm:"size(10)" json:"durable_years"`           // 预计使用年限
	BuyTime       time.Time `orm:"auto_now;type(datetime)" json:"buy_time"` // 购买时间
	Remark        string    `orm:"size(255)" json:"remark"`                 // 备注
}
