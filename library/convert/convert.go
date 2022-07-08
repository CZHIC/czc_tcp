package convert

import (
	"czc_tcp/library/logger"
	"encoding/json"
	"reflect"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

//Json Body转换为对应结构体
func JsonBody2Req(ctx *ghttp.Request, req interface{}) error {

	//原始数据
	content := ctx.GetBodyString()

	//xss转义
	bodyMap := gconv.Map(content)
	if bodyMap != nil {
		data := gconv.Map(bodyMap["data"])
		for k, v := range data {
			if gconv.String(reflect.TypeOf(v)) != "string" {
				continue
			}
			if k == "staff_info" {
				continue
			}
			data[k] = v
		}
		tmp := map[string]interface{}{
			"data": data,
		}
		bodyJson, err := json.Marshal(tmp)
		if err == nil {
			content = string(bodyJson)
		}
	}
	if content != "" {
		if err := json.Unmarshal([]byte(content), req); err != nil {
			return err
		}
	}
	return nil
}

//Json Body转换为对应结构体
func JsonBodyToStruct(ctx *ghttp.Request, req interface{}) error {
	content := ctx.GetBodyString()
	logger.Info("JsonBodyToStruct    content    ", content)
	if err := json.Unmarshal([]byte(content), req); err != nil {
		return err
	}
	return nil
}
