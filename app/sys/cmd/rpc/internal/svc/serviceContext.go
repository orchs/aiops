package svc

import (
	"aiops/app/sys/cmd/rpc/internal/config"
	"aiops/app/sys/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	//UserModel     model.UserModel
	BaseUserModel model.BaseUserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
		//UserModel:     model.NewUserModel(sqlConn, c.Cache),
		BaseUserModel: model.NewBaseUserModel(sqlConn, c.Cache),
	}
}
