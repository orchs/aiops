package user

import (
	"aiops/app/sys/cmd/api/internal/svc"
	"aiops/app/sys/cmd/api/internal/types"
	"aiops/app/sys/cmd/rpc/sys"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	userResp, err := l.svcCtx.SysRpc.EditUserInfo(l.ctx, &sys.EditUserInfoReq{
		Id: req.Id,
		User: &sys.User{
			Id:       req.UserInfo.Id,
			Nickname: req.UserInfo.Nickname,
			Mobile:   req.UserInfo.Mobile,
			Sex:      req.UserInfo.Sex,
			Avatar:   req.UserInfo.Avatar,
			Info:     req.UserInfo.Info,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	_ = copier.Copy(&resp, userResp)

	return
}
