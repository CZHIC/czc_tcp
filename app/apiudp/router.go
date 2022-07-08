package apiudp

import (
	"czc_tcp/app/model"
	"czc_tcp/library/logger"

	"github.com/gogf/gf/v2/net/gudp"
)

func Router(conn *gudp.Conn, msg *model.Msg) {
	switch msg.Act {
	case "Test":
		logger.Info("Udp 收到消息：", msg, " FROM:", conn.LocalAddr(), conn.RemoteAddr())
	case "heartbeat":
		logger.Info("Udp 收到消息：", msg, " FROM:", conn.LocalAddr(), conn.RemoteAddr())
	default:
		logger.Info("Udp 收到消息：", msg, " FROM:", conn.LocalAddr(), conn.RemoteAddr())
	}
}
