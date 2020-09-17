// 业主信息管理
// 记录业主基本信息、业主家庭成员、车辆、宠物等信息
package models

import "time"

const (
	TableEsMember = "es_member"
)

// 小区成员信息表
type EsMember struct {
	Id            int       `orm:"column(Id)" json:"id"`                           // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`                 // 所属小区编号，建议CM开头
	HouseCode     string    `orm:"size(50)" json:"house_code"`                     // 房产编号，建议HS开头
	Name          string    `orm:"size(50)" json:"name"`                           // 成员姓名
	IdentityId    string    `orm:"size(50)" json:"identity_id"`                    // 身份证号
	Tel           string    `orm:"size(50)" json:"tel"`                            // 联系方式
	Occupation    string    `orm:"size(255)" json:"occupation"`                    // 职业
	Birth         string    `orm:"size(20)" json:"birth"`                          // 出生日期
	Gender        int       `orm:"size(1)" json:"gender"`                          // 性别 0 女 1 男
	OwnerType     int       `orm:"size(1)" json:"owner_type"`                      // 成员类型 1 户主 2 家庭成员，3 租户
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"` // 添加时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"update_time"`     // 入住时间
	Remark        string    `orm:"size(255)" json:"remark"`                        // 备注
	Photo         string    `orm:"type(text)" json:"photo"`                        // 成员照片，拍照上传即可
}

// 车辆信息表
type EsVehicle struct {
	Id           int       `orm:"column(Id)"`                  // 主键
	MemberId     int       `orm:"size(11)"`                    // 家庭成员id
	Name         string    `orm:"size(255)"`                   // 车辆名称
	LicensePlate string    `orm:"size(50)"`                    // 车辆牌照
	Color        string    `orm:"size(50)"`                    // 车辆颜色
	Photo        string    `orm:"type(text)"`                  // 车辆照片，拍照上传即可
	Remark       string    `orm:"size(255)"`                   // 备注
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"` // 添加时间
	UpdateTime   time.Time `orm:"auto_now;type(datetime)"`     // 修改时间
}
