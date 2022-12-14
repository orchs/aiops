package svc

import (
	"aiops/app/order/cmd/rpc/order"
	"aiops/app/payment/cmd/api/internal/config"
	"aiops/app/payment/cmd/rpc/payment"
	"aiops/app/sys/cmd/rpc/sys"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	WxPayClient *core.Client

	PaymentRpc payment.Payment
	OrderRpc   order.Order
	SysRpc     sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		SysRpc:     sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
	}
}
