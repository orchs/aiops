package svc

import (
	"aiops/app/sys/cmd/api/internal/config"
	"aiops/app/sys/cmd/rpc/sys"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sys.Sys

	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		SysRpc: sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
	}
}
