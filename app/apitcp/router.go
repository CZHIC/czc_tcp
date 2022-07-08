package apitcp

import (
	"czc_tcp/app/model"
	"czc_tcp/library/logger"

	"github.com/gogf/gf/v2/net/gtcp"
)

func Router(conn *gtcp.Conn, msg *model.Msg) {
	switch msg.Act {
	case "hello":
		onClientHello(conn, msg)
	case "heartbeat":
		onClientHeartBeat(conn, msg)
	default:
		logger.Errorf("invalid message: %v", msg)
	}
}
