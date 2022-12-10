package svc

import (
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
	"aiops/app/mqueue/cmd/job/internal/config"
	"aiops/app/order/cmd/rpc/order"
	"aiops/app/sys/cmd/rpc/sys"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MiniProgram *miniprogram.MiniProgram

	OrderRpc order.Order
	SysRpc   sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		MiniProgram: newMiniprogramClient(c),
		OrderRpc:    order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		SysRpc:      sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
	}
}



