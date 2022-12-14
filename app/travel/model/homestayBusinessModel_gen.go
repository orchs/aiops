// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"aiops/common/globalkey"
	"aiops/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayBusinessFieldNames          = builder.RawFieldNames(&HomestayBusiness{})
	homestayBusinessRows                = strings.Join(homestayBusinessFieldNames, ",")
	homestayBusinessRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayBusinessFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayBusinessRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayBusinessFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheaiopsTravelHomestayBusinessIdPrefix     = "cache:aiopsTravel:homestayBusiness:id:"
	cacheaiopsTravelHomestayBusinessUserIdPrefix = "cache:aiopsTravel:homestayBusiness:userId:"
)

type (
	homestayBusinessModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayBusiness, error)
		FindOneByUserId(ctx context.Context, userId int64) (*HomestayBusiness, error)
		Update(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayBusinessModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayBusiness struct {
		Id          int64     `db:"id"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
		DeleteTime  time.Time `db:"delete_time"`
		DelState    int64     `db:"del_state"`
		Title       string    `db:"title"`        // 店铺名称
		UserId      int64     `db:"user_id"`      // 关联的用户id
		Info        string    `db:"info"`         // 店铺介绍
		BossInfo    string    `db:"boss_info"`    // 房东介绍
		LicenseFron string    `db:"license_fron"` // 营业执照正面
		LicenseBack string    `db:"license_back"` // 营业执照背面
		RowState    int64     `db:"row_state"`    // 0:禁止营业 1:正常营业
		Star        float64   `db:"star"`         // 店铺整体评价，冗余
		Tags        string    `db:"tags"`         // 每个店家一个标签，自己编辑
		Cover       string    `db:"cover"`        // 封面图
		HeaderImg   string    `db:"header_img"`   // 店招门头图片
		Version     int64     `db:"version"`      // 版本号
	}
)

func newHomestayBusinessModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayBusinessModel {
	return &defaultHomestayBusinessModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_business`",
	}
}

func (m *defaultHomestayBusinessModel) Insert(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	aiopsTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, data.Id)
	aiopsTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayBusinessRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version)
	}, aiopsTravelHomestayBusinessIdKey, aiopsTravelHomestayBusinessUserIdKey)
}

func (m *defaultHomestayBusinessModel) FindOne(ctx context.Context, id int64) (*HomestayBusiness, error) {
	aiopsTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, id)
	var resp HomestayBusiness
	err := m.QueryRowCtx(ctx, &resp, aiopsTravelHomestayBusinessIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindOneByUserId(ctx context.Context, userId int64) (*HomestayBusiness, error) {
	aiopsTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessUserIdPrefix, userId)
	var resp HomestayBusiness
	err := m.QueryRowIndexCtx(ctx, &resp, aiopsTravelHomestayBusinessUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) Update(ctx context.Context, session sqlx.Session, data *HomestayBusiness) (sql.Result, error) {
	aiopsTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, data.Id)
	aiopsTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayBusinessRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version, data.Id)
	}, aiopsTravelHomestayBusinessIdKey, aiopsTravelHomestayBusinessUserIdKey)
}

func (m *defaultHomestayBusinessModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	aiopsTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, data.Id)
	aiopsTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessUserIdPrefix, data.UserId)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayBusinessRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version, data.Id, oldVersion)
	}, aiopsTravelHomestayBusinessIdKey, aiopsTravelHomestayBusinessUserIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultHomestayBusinessModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	aiopsTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, id)
	aiopsTravelHomestayBusinessUserIdKey := fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, aiopsTravelHomestayBusinessIdKey, aiopsTravelHomestayBusinessUserIdKey)
	return err
}

func (m *defaultHomestayBusinessModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheaiopsTravelHomestayBusinessIdPrefix, primary)
}
func (m *defaultHomestayBusinessModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayBusinessRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayBusinessModel) tableName() string {
	return m.table
}
