package main

import (
	"czc_tcp/app/apitcp"
	"czc_tcp/app/apiudp"
	_ "czc_tcp/boot"
	"czc_tcp/library/logger"
	"czc_tcp/library/pubFuc"
	_ "czc_tcp/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
)

func main() {
	go g.Server().Run()

	//udp
	go gudp.NewServer("0.0.0.0:8899", func(conn *gudp.Conn) {
		logger.Println("UDP服务启动：", conn.LocalAddr(), conn.RemoteAddr())
		defer conn.Close()
		for {
			msg, err := pubFuc.UdpRecvPkg(conn)
			if err == nil {
				if err := pubFuc.UdpSendPkg(conn, msg); err != nil {
					g.Log().Error(err)
				}
			}
			if err != nil {
				g.Log().Error(err)
			}
			apiudp.Router(conn, msg)
		}
	}).Run()

	// tcp
	gtcp.NewServer("0.0.0.0:8090", func(conn *gtcp.Conn) {
		defer conn.Close()
		logger.Println("TCP服务启动：", conn.LocalAddr(), conn.RemoteAddr())
		for {
			msg, err := pubFuc.RecvPkg(conn)
			if err != nil {
				if err.Error() == "EOF" {
					glog.Println("client closed")
				}
				break
			}
			apitcp.Router(conn, msg)
		}

	}).Run()

}
