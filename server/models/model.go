package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// 自动注册模型
func init() {
	// 1. 注册数据驱动, mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		logs.Error("register Driver sqlite error, error message: ", err)
	}

	// 2. 注册数据库, ORM必须注册一个别名为 default 的数据库，作为默认使用
	/*if err := orm.RegisterDataBase("default", "sqlite3", "./test.db"); err != nil {
		logs.Error("register DataBase error, error message: ", err)
	}*/
	if err := orm.RegisterDataBase("default", "mysql",
		`root:root@/property_management?charset=utf8`); err != nil {
		logs.Error("register DataBase error, error message: ", err)
	}

	// 3. 注册模型
	orm.RegisterModel(
		new(User),              // 登录用户表
		new(EsDevice),          // 资产设备信息表
		new(EsExpensesProject), // 收费项目信息表
		new(EsExpenses),        // 费用信息表
		new(EsCommunity),       // 小区信息表
		new(EsBuilding),        // 栋数信息表
		new(EsHouse),           // 房产信息表
		new(EsMember),          // 小区成员信息表
		new(EsVehicle),         // 车辆信息表
		new(EsParkingSpace),    // 停车位基本信息表
		new(EsParkingSpaceUse), // 停车位使用记录表
		new(EsDuty),            // 物业值班人员表
		new(Admin),             // 管理员表
	)

	// 4. 自动创建表 参数二为是否开启创建表   参数三是否更新表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error("run sync db error, error message: ", err)
	}
}

func TableName(name string) string {
	return name
}
