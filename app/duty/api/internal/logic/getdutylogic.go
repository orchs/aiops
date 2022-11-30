package logic

import (
	"aiops/app/user/rpc/userclient"
	"context"
	"errors"

	"aiops/app/duty/api/internal/svc"
	"aiops/app/duty/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDutyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDutyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDutyLogic {
	return &GetDutyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDutyLogic) GetDuty(req *types.DutyReq) (resp *types.DutyReply, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: 1,
	})
	if err != nil {
		return nil, err
	}

	if user.Username != "wk" {
		return nil, errors.New("用户不存在")
	}

	logx.Infof("userId: %v", l.ctx.Value("userId"))

	return &types.DutyReply{
		Id:   req.Id,
		Name: "test duty",
	}, nil
}
