package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BaseUserModel = (*customBaseUserModel)(nil)

type (
	// BaseUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBaseUserModel.
	BaseUserModel interface {
		baseUserModel
	}

	customBaseUserModel struct {
		*defaultBaseUserModel
	}
)

// NewBaseUserModel returns a model for the database table.
func NewBaseUserModel(conn sqlx.SqlConn, c cache.CacheConf) BaseUserModel {
	return &customBaseUserModel{
		defaultBaseUserModel: newBaseUserModel(conn, c),
	}
}
