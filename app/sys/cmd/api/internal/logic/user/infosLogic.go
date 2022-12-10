package user

import (
	"context"

	"aiops/app/sys/cmd/api/internal/svc"
	"aiops/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfosLogic {
	return &InfosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfosLogic) Infos() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
