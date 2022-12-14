// Code generated by goctl. DO NOT EDIT!

package model

import (
	"aiops/deploy/script/mysql/genModel"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	baseUserFieldNames          = builder.RawFieldNames(&BaseUser{})
	baseUserRows                = strings.Join(baseUserFieldNames, ",")
	baseUserRowsExpectAutoSet   = strings.Join(stringx.Remove(baseUserFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	baseUserRowsWithPlaceHolder = strings.Join(stringx.Remove(baseUserFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"

	cacheAiopsSysBaseUserIdPrefix       = "cache:aiopsSys:baseUser:id:"
	cacheAiopsSysBaseUserMobilePrefix   = "cache:aiopsSys:baseUser:mobile:"
	cacheAiopsSysBaseUserUsernamePrefix = "cache:aiopsSys:baseUser:username:"
)

type (
	baseUserModel interface {
		Insert(ctx context.Context, data *BaseUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BaseUser, error)
		FindOneByMobile(ctx context.Context, mobile string) (*BaseUser, error)
		FindOneByUsername(ctx context.Context, username string) (*BaseUser, error)
		Update(ctx context.Context, data *BaseUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBaseUserModel struct {
		sqlc.CachedConn
		table string
	}

	BaseUser struct {
		Id       int64        `db:"id"`
		CreateBy string       `db:"create_by"`
		CreateAt time.Time    `db:"create_at"`
		UpdateBy string       `db:"update_by"`
		UpdateAt time.Time    `db:"update_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
		Username string       `db:"username"`
		Version  int64        `db:"version"` // 版本号
		Mobile   string       `db:"mobile"`
		Password string       `db:"password"`
		Nickname string       `db:"nickname"`
		Sex      int64        `db:"sex"` // 性别 0:男 1:女
		Avatar   string       `db:"avatar"`
		Info     string       `db:"info"`
	}
)

func newBaseUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBaseUserModel {
	return &defaultBaseUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`base_user`",
	}
}

func (m *defaultBaseUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	aiopsSysBaseUserIdKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserIdPrefix, id)
	aiopsSysBaseUserMobileKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserMobilePrefix, data.Mobile)
	aiopsSysBaseUserUsernameKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, aiopsSysBaseUserIdKey, aiopsSysBaseUserMobileKey, aiopsSysBaseUserUsernameKey)
	return err
}

func (m *defaultBaseUserModel) FindOne(ctx context.Context, id int64) (*BaseUser, error) {
	aiopsSysBaseUserIdKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserIdPrefix, id)
	var resp BaseUser
	err := m.QueryRowCtx(ctx, &resp, aiopsSysBaseUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBaseUserModel) FindOneByMobile(ctx context.Context, mobile string) (*BaseUser, error) {
	aiopsSysBaseUserMobileKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserMobilePrefix, mobile)
	var resp BaseUser
	err := m.QueryRowIndexCtx(ctx, &resp, aiopsSysBaseUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", baseUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBaseUserModel) FindOneByUsername(ctx context.Context, username string) (*BaseUser, error) {
	aiopsSysBaseUserUsernameKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserUsernamePrefix, username)
	var resp BaseUser
	err := m.QueryRowIndexCtx(ctx, &resp, aiopsSysBaseUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", baseUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBaseUserModel) Insert(ctx context.Context, data *BaseUser) (sql.Result, error) {
	aiopsSysBaseUserIdKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserIdPrefix, data.Id)
	aiopsSysBaseUserMobileKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserMobilePrefix, data.Mobile)
	aiopsSysBaseUserUsernameKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, baseUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateBy, data.UpdateBy, data.DeleteAt, data.Username, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
	}, aiopsSysBaseUserIdKey, aiopsSysBaseUserMobileKey, aiopsSysBaseUserUsernameKey)
	return ret, err
}

func (m *defaultBaseUserModel) Update(ctx context.Context, newData *BaseUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	aiopsSysBaseUserIdKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserIdPrefix, data.Id)
	aiopsSysBaseUserMobileKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserMobilePrefix, data.Mobile)
	aiopsSysBaseUserUsernameKey := fmt.Sprintf("%s%v", cacheAiopsSysBaseUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, baseUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.CreateBy, newData.UpdateBy, newData.DeleteAt, newData.Username, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.Id)
	}, aiopsSysBaseUserIdKey, aiopsSysBaseUserMobileKey, aiopsSysBaseUserUsernameKey)
	return err
}

func (m *defaultBaseUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAiopsSysBaseUserIdPrefix, primary)
}

func (m *defaultBaseUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBaseUserModel) tableName() string {
	return m.table
}
