package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BaseRoleModel = (*customBaseRoleModel)(nil)

type (
	// BaseRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBaseRoleModel.
	BaseRoleModel interface {
		baseRoleModel
	}

	customBaseRoleModel struct {
		*defaultBaseRoleModel
	}
)

// NewBaseRoleModel returns a model for the database table.
func NewBaseRoleModel(conn sqlx.SqlConn, c cache.CacheConf) BaseRoleModel {
	return &customBaseRoleModel{
		defaultBaseRoleModel: newBaseRoleModel(conn, c),
	}
}
