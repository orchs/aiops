package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelUserTeamModel = (*customRelUserTeamModel)(nil)

type (
	// RelUserTeamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelUserTeamModel.
	RelUserTeamModel interface {
		relUserTeamModel
	}

	customRelUserTeamModel struct {
		*defaultRelUserTeamModel
	}
)

// NewRelUserTeamModel returns a model for the database table.
func NewRelUserTeamModel(conn sqlx.SqlConn, c cache.CacheConf) RelUserTeamModel {
	return &customRelUserTeamModel{
		defaultRelUserTeamModel: newRelUserTeamModel(conn, c),
	}
}
