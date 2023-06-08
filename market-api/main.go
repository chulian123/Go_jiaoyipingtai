package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"market-api/internal/config"
	"market-api/internal/handler"
	"market-api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
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
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,x-auth-token")
	}, nil, "http://localhost:8080"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	routers := handler.NewRouters(server, c.Prefix)
	handler.RegisterHandlers(routers, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
