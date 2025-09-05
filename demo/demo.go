package main

import (
	"flag"
	"fmt"

	"demo/internal/config"
	"demo/internal/handler"
	"demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 1) 加载配置：读取 YAML 并反序列化到 config.Config
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 2) 初始化 REST Server（go-zero）
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 3) 构建 ServiceContext（初始化 DB 等依赖）并注册路由
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
