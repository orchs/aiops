package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BaseTeamModel = (*customBaseTeamModel)(nil)

type (
	// BaseTeamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBaseTeamModel.
	BaseTeamModel interface {
		baseTeamModel
	}

	customBaseTeamModel struct {
		*defaultBaseTeamModel
	}
)

// NewBaseTeamModel returns a model for the database table.
func NewBaseTeamModel(conn sqlx.SqlConn, c cache.CacheConf) BaseTeamModel {
	return &customBaseTeamModel{
		defaultBaseTeamModel: newBaseTeamModel(conn, c),
	}
}
