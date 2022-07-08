package apitcp

import (
	"czc_tcp/app/model"
	"czc_tcp/library/logger"
	"czc_tcp/library/pubFuc"

	"github.com/gogf/gf/v2/net/gtcp"
)

func onClientHello(conn *gtcp.Conn, msg *model.Msg) {
	logger.Infof("hello message from [%s]: %s", conn.RemoteAddr().String(), msg.Data)
	pubFuc.SendPkg(conn, msg.Act, "Nice to meet you!")
}

func onClientHeartBeat(conn *gtcp.Conn, msg *model.Msg) {
	logger.Infof("heartbeat from [%s]", conn.RemoteAddr().String())
}
