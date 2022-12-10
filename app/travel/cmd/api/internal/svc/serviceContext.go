package svc

import (
	"aiops/app/travel/cmd/api/internal/config"
	"aiops/app/travel/cmd/rpc/travel"
	"aiops/app/travel/model"
	"aiops/app/sys/cmd/rpc/sys"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	//local
	Config config.Config

	//rpc
	SysRpc    sys.Sys
	TravelRpc travel.Travel

	//model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		SysRpc:    sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
		TravelRpc: travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),

		HomestayModel:         model.NewHomestayModel(sqlConn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlConn, c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(sqlConn, c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(sqlConn, c.Cache),
	}
}
