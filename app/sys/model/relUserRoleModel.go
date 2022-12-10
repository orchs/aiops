package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelUserRoleModel = (*customRelUserRoleModel)(nil)

type (
	// RelUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelUserRoleModel.
	RelUserRoleModel interface {
		relUserRoleModel
	}

	customRelUserRoleModel struct {
		*defaultRelUserRoleModel
	}
)

// NewRelUserRoleModel returns a model for the database table.
func NewRelUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) RelUserRoleModel {
	return &customRelUserRoleModel{
		defaultRelUserRoleModel: newRelUserRoleModel(conn, c),
	}
}
