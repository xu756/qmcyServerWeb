package main

import (
	"flag"
	"fmt"
	"github.com/xu756/qmcy/pb"

	"github.com/xu756/qmcy/app/public/rpc/internal/config"
	"github.com/xu756/qmcy/app/public/rpc/internal/server"
	"github.com/xu756/qmcy/app/public/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/public.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPublicServer(grpcServer, server.NewPublicServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("【 public-rpc 】  rpc server at %s...\n", c.ListenOn)
	s.Start()
}
