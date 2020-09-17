package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

// 响应数据格式
type Response struct {
	Code int         `json:"code"` // 响应值
	Data interface{} `json:"data"` // 回执数据
}

// 返回查询数据格式
type ResponseTable struct {
	Total int         `json:"total"` // 总数
	Items interface{} `json:"items"` // 查询结果
}

// 结构体转json
func StructToJSON(s interface{}) (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		logs.Error("struct to json error, error message: ", err)
		return "", err
	}
	return string(data), nil
}
