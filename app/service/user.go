package service

import (
	"czc_tcp/app/dao"
	"czc_tcp/app/model"
	"czc_tcp/library/logger"

	"github.com/gogf/gf/os/gtime"
)

type UserReq struct {
	Phone int64  `json:"phone"`
	Name  string `json:"name"`
}

// 判断是否是高层 svp或BDIC委员
func Login(req *UserReq) (*model.User, error) {
	ret, err := dao.GetUserByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		ret = &model.User{}
	}
	logger.Println("------ret-----", ret)
	ret.CreateAt = gtime.Now()
	ret.Phone = req.Phone
	ret.Name = req.Name
	err = dao.UpdateOrInsert(ret)
	if err != nil {
		return nil, err
	}
	result, err := dao.GetUserByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	return result, nil
}
