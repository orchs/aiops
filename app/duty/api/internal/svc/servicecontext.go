package svc

import (
	"aiops/app/duty/api/internal/config"
	"aiops/app/duty/api/internal/middleware"
	"aiops/app/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
