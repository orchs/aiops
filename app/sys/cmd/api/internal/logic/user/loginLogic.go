package user

import (
	"context"

	"aiops/app/sys/cmd/api/internal/svc"
	"aiops/app/sys/cmd/api/internal/types"
	"aiops/app/sys/cmd/rpc/sys"
	"aiops/app/sys/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	loginResp, err := l.svcCtx.SysRpc.Login(l.ctx, &sys.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)

	return &resp, nil
}
