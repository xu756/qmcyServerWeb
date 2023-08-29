package svc

import (
	"github.com/xu756/qmcy/app/admin/api/internal/config"
	"github.com/xu756/qmcy/app/admin/api/internal/middleware"
	"github.com/xu756/qmcy/app/admin/rpc/admin"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Auth     rest.Middleware
	AdminRpc admin.Admin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Auth:     middleware.NewAuthMiddleware(c).Handle,
		AdminRpc: admin.NewAdmin(zrpc.MustNewClient(c.Admin)),
	}
}
