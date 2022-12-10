// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	relUserRoleFieldNames          = builder.RawFieldNames(&RelUserRole{})
	relUserRoleRows                = strings.Join(relUserRoleFieldNames, ",")
	relUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(relUserRoleFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	relUserRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(relUserRoleFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"

	cacheAiopsSysRelUserRoleIdPrefix           = "cache:aiopsSys:relUserRole:id:"
	cacheAiopsSysRelUserRoleRoleIdUserIdPrefix = "cache:aiopsSys:relUserRole:roleId:userId:"
	cacheAiopsSysRelUserRoleUserIdRoleIdPrefix = "cache:aiopsSys:relUserRole:userId:roleId:"
)

type (
	relUserRoleModel interface {
		Insert(ctx context.Context, data *RelUserRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RelUserRole, error)
		FindOneByRoleIdUserId(ctx context.Context, roleId int64, userId int64) (*RelUserRole, error)
		FindOneByUserIdRoleId(ctx context.Context, userId int64, roleId int64) (*RelUserRole, error)
		Update(ctx context.Context, data *RelUserRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRelUserRoleModel struct {
		sqlc.CachedConn
		table string
	}

	RelUserRole struct {
		Id     int64 `db:"id"`
		UserId int64 `db:"user_id"`
		RoleId int64 `db:"role_id"`
	}
)

func newRelUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRelUserRoleModel {
	return &defaultRelUserRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`rel_user_role`",
	}
}

func (m *defaultRelUserRoleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	aiopsSysRelUserRoleIdKey := fmt.Sprintf("%s%v", cacheAiopsSysRelUserRoleIdPrefix, id)
	aiopsSysRelUserRoleRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleRoleIdUserIdPrefix, data.RoleId, data.UserId)
	aiopsSysRelUserRoleUserIdRoleIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleUserIdRoleIdPrefix, data.UserId, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, aiopsSysRelUserRoleIdKey, aiopsSysRelUserRoleRoleIdUserIdKey, aiopsSysRelUserRoleUserIdRoleIdKey)
	return err
}

func (m *defaultRelUserRoleModel) FindOne(ctx context.Context, id int64) (*RelUserRole, error) {
	aiopsSysRelUserRoleIdKey := fmt.Sprintf("%s%v", cacheAiopsSysRelUserRoleIdPrefix, id)
	var resp RelUserRole
	err := m.QueryRowCtx(ctx, &resp, aiopsSysRelUserRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relUserRoleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultRelUserRoleModel) FindOneByRoleIdUserId(ctx context.Context, roleId int64, userId int64) (*RelUserRole, error) {
	aiopsSysRelUserRoleRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleRoleIdUserIdPrefix, roleId, userId)
	var resp RelUserRole
	err := m.QueryRowIndexCtx(ctx, &resp, aiopsSysRelUserRoleRoleIdUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `role_id` = ? and `user_id` = ? limit 1", relUserRoleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, roleId, userId); err != nil {
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

func (m *defaultRelUserRoleModel) FindOneByUserIdRoleId(ctx context.Context, userId int64, roleId int64) (*RelUserRole, error) {
	aiopsSysRelUserRoleUserIdRoleIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleUserIdRoleIdPrefix, userId, roleId)
	var resp RelUserRole
	err := m.QueryRowIndexCtx(ctx, &resp, aiopsSysRelUserRoleUserIdRoleIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `role_id` = ? limit 1", relUserRoleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, roleId); err != nil {
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

func (m *defaultRelUserRoleModel) Insert(ctx context.Context, data *RelUserRole) (sql.Result, error) {
	aiopsSysRelUserRoleIdKey := fmt.Sprintf("%s%v", cacheAiopsSysRelUserRoleIdPrefix, data.Id)
	aiopsSysRelUserRoleRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleRoleIdUserIdPrefix, data.RoleId, data.UserId)
	aiopsSysRelUserRoleUserIdRoleIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleUserIdRoleIdPrefix, data.UserId, data.RoleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, relUserRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.RoleId)
	}, aiopsSysRelUserRoleIdKey, aiopsSysRelUserRoleRoleIdUserIdKey, aiopsSysRelUserRoleUserIdRoleIdKey)
	return ret, err
}

func (m *defaultRelUserRoleModel) Update(ctx context.Context, newData *RelUserRole) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	aiopsSysRelUserRoleIdKey := fmt.Sprintf("%s%v", cacheAiopsSysRelUserRoleIdPrefix, data.Id)
	aiopsSysRelUserRoleRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleRoleIdUserIdPrefix, data.RoleId, data.UserId)
	aiopsSysRelUserRoleUserIdRoleIdKey := fmt.Sprintf("%s%v:%v", cacheAiopsSysRelUserRoleUserIdRoleIdPrefix, data.UserId, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, relUserRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.RoleId, newData.Id)
	}, aiopsSysRelUserRoleIdKey, aiopsSysRelUserRoleRoleIdUserIdKey, aiopsSysRelUserRoleUserIdRoleIdKey)
	return err
}

func (m *defaultRelUserRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAiopsSysRelUserRoleIdPrefix, primary)
}

func (m *defaultRelUserRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relUserRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRelUserRoleModel) tableName() string {
	return m.table
}