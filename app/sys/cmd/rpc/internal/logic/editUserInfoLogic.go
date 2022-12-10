package logic

import (
	"aiops/app/sys/cmd/rpc/sys"
	"aiops/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"aiops/app/sys/cmd/rpc/internal/svc"
	"aiops/app/sys/cmd/rpc/pb"
	"aiops/app/sys/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserInfoLogic {
	return &EditUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserInfoLogic) EditUserInfo(in *pb.EditUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.BaseUserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据id查询用户信息失败，user_id:%d,err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "user_id:%d", in.Id)
	}

	if user.Id != in.User.Id {
		return nil, errors.Wrapf(
			xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR),
			"要修改的用户id和url中的参数不一致，请确认！",
		)
	}

	user.Nickname = in.User.Nickname
	user.Sex = in.User.Sex
	user.Info = in.User.Info
	user.Mobile = in.User.Mobile
	l.svcCtx.BaseUserModel.Update(l.ctx, user)

	var respUser sys.User
	_ = copier.Copy(&respUser, user)

	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
