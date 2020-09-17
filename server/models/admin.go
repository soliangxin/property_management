// 权限管理
// 权限管理 详细的操作权限控制，不仅仅是菜单级别的控制，同时还可以精确到界面上的操作按钮。
// 系统内置了以下4种角色： 超级管理员、小区管理员、小区普通员工、业主 每一种角色都有不同的操作权限。
// 同时也可以根据需要，自定义多种角色，然后为其分配相应的操作权限。
package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 登录用户表, 授权登录用户相关
type User struct {
	Id         int       `orm:"column(Id)"`                  // 主键
	Username   string    `orm:"size(16);unique"`             // 用户名
	Password   string    `orm:"size(32)"`                    // 密码
	Sex        int       `orm:"size(1)"`                     // 性别
	Email      string    `orm:"size(50);unique"`             // 邮箱
	Phone      string    `orm:"size(50);unique"`             // 联系电话
	IsAdmin    int       `orm:"size(1);default(0)"`          // 是否是管理员, 0 否  2 是
	Status     int       `orm:"size(1);default(0)"`          // 状态, 0 禁用 2 启用
	LoginTime  time.Time `orm:"auto_now;type(datetime)"`     // 登录时间
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 添加时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 修改时间
}

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	Delete     int
	CreateTime int64
	UpdateTime int64
}


const TableAdminUser = "admin_user"

func (a *Admin) TableName() string {
	return "admin_user"
}

func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableAdminUser).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm().QueryTable(TableAdminUser)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	_, _ = query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	err := orm.NewOrm().QueryTable(&Admin{}).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
