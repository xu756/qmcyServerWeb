package svc

import (
	"github.com/xu756/qmcy/app/user/api/internal/config"
	"github.com/xu756/qmcy/app/user/api/internal/middleware"
	"github.com/xu756/qmcy/app/user/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware(c).Handle,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.User)),
	}
}
