package request

import (
	"czc_tcp/library/convert"
	"czc_tcp/library/logger"

	"github.com/gogf/gf/net/ghttp"
)

// request结构体
type BaseRequest struct {
	Data interface{} `json:"data"`
}

//获取请求参数中的data信息
func GetStructData(ctx *ghttp.Request, data interface{}) (interface{}, error) {
	req := new(BaseRequest)
	req.Data = data
	if err := convert.JsonBody2Req(ctx, req); err != nil {
		logger.Errorf("----GetStructData---error----", err.Error())
		return nil, err
	}
	return req.Data, nil
}

//获取请求参数中的data信息
func GetJsonStructData(ctx *ghttp.Request, data interface{}) (interface{}, error) {
	if err := convert.JsonBodyToStruct(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Staff结构体
type StaffInfo struct {
	StaffId   int64  `json:"staff_id"`
	StaffName string `json:"staff_name"`
}
