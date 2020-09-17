// 房产信息管理
// 记录每个住户的基本信息，包括户主，房间数，单元信息，楼层
package models

import "time"

const (
	TableEsCommunity = "es_community" // 小区信息表
	TableEsHouse     = "es_house"
)

// 小区信息表
type EsCommunity struct {
	Id             int       `orm:"" json:"id"`                                  // 主键
	CommunityNum   string    `orm:"size(50);unique" json:"community_num"`        // 小区编号，建议CM开头
	Name           string    `orm:"size(255)" json:"name"`                       // 小区名称
	Introduction   string    `orm:"type(text)" json:"introduction"`              // 简介
	Address        string    `orm:"size(255)" json:"address"`                    // 坐落地址
	Area           float64   `orm:"digits(15);decimals(2)" json:"area"`          // 占地面积，单位：平米
	Developer      string    `orm:"size(255)" json:"developer"`                  // 开发商名称
	Estate         string    `orm:"size(255)" json:"estate"`                     // 物业公司名称
	GreeningRate   float64   `orm:"digits(10);decimals(2)" json:"greening_rate"` // 绿化率，单位：百分比
	TotalBuilding  int       `orm:"size(11)" json:"total_building"`              // 总栋数
	TotalOwner     int       `orm:"size(11)" json:"total_owner"`                 // 总户数
	BuildTime      time.Time `orm:"type(date)" json:"build_time"`                // 修建时间
	CompletionTime time.Time `orm:"type(date)" json:"completion_time"`           // 竣工时间
	UpdateTime     time.Time `orm:"auto_now;type(datetime)" json:"update_time"`  // 修改时间
}

// 栋数信息表
type EsBuilding struct {
	Id            int       `orm:"column(Id)"`                  // 主键
	CommunityCode string    `orm:"size(50)"`                    // 所属小区编号，建议CM开头
	Code          string    `orm:"size(50);unique"`             // 栋数编号，建议BD开头
	Name          string    `orm:"size(255)"`                   // 栋数名称
	House         int       `orm:"size(10)"`                    // 总户数
	Desc          string    `orm:"size(255)"`                   // 描述
	Lift          int       `orm:"size(10)"`                    // 电梯数
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"` // 添加时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`     // 修改时间
}

// 房产信息表
type EsHouse struct {
	Id            int       `orm:"column(Id)" json:"id"`                          // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`                // 所属小区编号，建议CM开头
	BuildingCode  string    `orm:"size(50)" json:"building_code"`                 // 栋数编号，建议BD开头
	HouseCode     string    `orm:"size(50);unique" json:"house_code"`             // 房产编号，建议HS开头
	Name          string    `orm:"size(255)" json:"name"`                         // 房产名称
	OwnerName     string    `orm:"size(100)" json:"owner_name"`                   // 户主名称
	OwnerTel      string    `orm:"size(100)" json:"owner_tel"`                    // 户主联系方式
	Rooms         int       `orm:"size(10)" json:"rooms"`                         // 房间数
	Unit          string    `orm:"size(50)" json:"unit"`                          // 单元信息
	Floor         int       `orm:"size(10)" json:"floor"`                         // 楼层信息
	HouseDesc     string    `orm:"size(255)" json:"house_desc"`                   // 房产描述
	EnterTime     time.Time `orm:"auto_now_add;type(datetime)" json:"enter_time"` // 入住时间
}
