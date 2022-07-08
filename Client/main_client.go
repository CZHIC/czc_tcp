package main

import (
	"czc_tcp/app/model"
	"czc_tcp/library/logger"
	"czc_tcp/library/pubFuc"
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/gogf/gf/v2/os/gtimer"
)

func main() {
	// var (
	// 	ctx = gctx.New()
	// )
	// TcpClient(ctx)

	UdpClient()

}

func TcpClient(ctx context.Context) {
	conn, err := gtcp.NewConn("127.0.0.1:8090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 心跳消息
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		if err := pubFuc.SendPkg(conn, "heartbeat"); err != nil {
			panic(err)
		}
	})
	// 测试消息, 3秒后向服务端发送hello消息
	gtimer.SetTimeout(ctx, 3*time.Second, func(ctx context.Context) {
		if err := pubFuc.SendPkg(conn, "hello", "My name's John!"); err != nil {
			panic(err)
		}
	})
	for {
		msg, err := pubFuc.RecvPkg(conn)
		if err != nil {
			if err.Error() == "EOF" {
				glog.Println("server closed")
			}
			break
		}
		switch msg.Act {
		case "hello":
			onServerHello(conn, msg)
		case "doexit":
			onServerDoExit(conn, msg)
		case "heartbeat":
			onServerHeartBeat(conn, msg)
		default:
			glog.Errorf("invalid message: %v", msg)
		}
	}
}

func UdpClient() {

	msg := model.Msg{
		Act:  "Test",
		Data: "hello world",
	}
	sendMsg, _ := json.Marshal(msg)
	// Client
	for {
		if conn, err := gudp.NewConn("9.134.189.214:11111"); err == nil {
			fmt.Println("服务启动：", conn.LocalAddr(), conn.RemoteAddr())
			if b, err := conn.SendRecv(sendMsg, -1); err == nil {
				fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
			} else {
				logger.Error(err)
			}
			conn.Close()
		} else {

			logger.Error(err)
		}
		time.Sleep(time.Second)
	}

}

func onServerHello(conn *gtcp.Conn, msg *model.Msg) {
	glog.Printf("hello response message from [%s]: %s", conn.RemoteAddr().String(), msg.Data)
}

func onServerHeartBeat(conn *gtcp.Conn, msg *model.Msg) {
	glog.Printf("heartbeat from [%s]", conn.RemoteAddr().String())
}

func onServerDoExit(conn *gtcp.Conn, msg *model.Msg) {
	glog.Printf("exit command from [%s]", conn.RemoteAddr().String())
	conn.Close()
}
