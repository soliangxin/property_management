package models

import "time"

const (
	TableEsDuty = "es_duty" // 物业员工值班信息表
)

// 物业员工值班信息表
type EsDuty struct {
	Id            int       `orm:"" json:"id"`                                    // 主键
	CommunityCode string    `orm:"size(50)" json:"community_code"`                // 所属小区编号，建议CM开头
	Name          string    `orm:"size(255)" json:"name"`                         // 值班人名称，多个值班人用英文逗号隔开
	StartTime     time.Time `orm:"auto_now_add;type(datetime)" json:"start_time"` // 值班开始时间
	EndTime       time.Time `orm:"auto_now;type(datetime)" json:"end_time"`       // 值班结束时间
	Remark        string    `orm:"size(255)" json:"remark"`                       // 备注
}
