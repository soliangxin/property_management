// 车位管理
// 记录车位基本信息（车位编号） 记录车位的使用情况（租或买），费用等信息
package models

import "time"

const (
	TableEsParkingSpace = "es_parking_space"
)

// 停车位基本信息表
type EsParkingSpace struct {
	Id            int       `orm:"" json:"id"`                                     // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`                 // 所属小区编号，建议CM开头
	Code          string    `orm:"size(50);unique" json:"code"`                    // 车位编号，建议PK开头
	Name          string    `orm:"size(255)" json:"name"`                          // 车位名称
	Status        int       `orm:"size(1);default(0)" json:"status"`               // 状态 0 闲置中 1 使用中
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"` // 添加时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"update_time"`     // 添加时间
}

// 停车位使用记录表
type EsParkingSpaceUse struct {
	Id            int       `orm:""`                            // 主键
	CommunityCode string    `orm:"size(50)"`                    // 所属小区编号，建议CM开头
	PkCode        string    `orm:"size(50)"`                    // 车位编号，建议PK开头
	LicensePlate  string    `orm:"size(50)"`                    // 车辆牌照
	Owner         string    `orm:"size(100)"`                   // 车辆所有人
	Tel           string    `orm:"size(50)"`                    // 联系电话
	BeginTime     time.Time `orm:"auto_now_add;type(datetime)"` // 开始时间
	EndTime       time.Time `orm:"auto_now_add;type(datetime)"` // 开始时间
	Status        int       `orm:"size(1);default(0)"`          // 使用性质 1 租 2 买
	Cost          float64   `orm:"digits(10);decimals(2)"`      // 费用
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"` // 添加时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`     // 修改时间
}
