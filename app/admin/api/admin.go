package main

import (
	"flag"
	"fmt"

	"github.com/xu756/qmcy/app/admin/api/internal/config"
	"github.com/xu756/qmcy/app/admin/api/internal/handler"
	"github.com/xu756/qmcy/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("【 admin-api 】  at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
