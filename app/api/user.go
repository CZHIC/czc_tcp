package api

import (
	"czc_tcp/app/service"
	"czc_tcp/library/request"
	"czc_tcp/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var UserApi = new(UserAPI)

type UserAPI struct{}

// @summary 定时清理离职人员权限
// @tags    Test-http
// @produce json
// @router  /user/login [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *UserAPI) Login(r *ghttp.Request) {
	reqData, _ := request.GetStructData(r, new(service.UserReq))

	req := reqData.(*service.UserReq)
	ret, err := service.Login(req)
	if err != nil {
		response.JsonExit(r, 1, "error", err.Error())
	}
	response.JsonExit(r, 0, "ok", ret)
}
