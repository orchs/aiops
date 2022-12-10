package svc

import (
	"aiops/app/order/cmd/mq/internal/config"
	"aiops/app/order/cmd/rpc/order"
	"aiops/app/sys/cmd/rpc/sys"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc order.Order
	SysRpc   sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		SysRpc:   sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
	}
}
