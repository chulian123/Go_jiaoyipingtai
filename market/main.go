package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"market/internal/config"
	"market/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()
	//日志的格式替换
	logx.MustSetup(logx.LogConf{
		Stat:     false,
		Encoding: "plain",
	})

	var c config.Config
	conf.MustLoad(*configFile, &c)
	_ = svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		//	register.RegisterRegisterServer(grpcServer, server.NewRegisterServer(ctx))
		//	login.RegisterLoginServer(grpcServer, server.NewLoginServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
