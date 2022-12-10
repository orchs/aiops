package homestayBussiness

import (
	"context"

	"aiops/app/travel/cmd/api/internal/svc"
	"aiops/app/travel/cmd/api/internal/types"
	"aiops/app/sys/cmd/rpc/sys"
	"aiops/app/sys/model"
	"aiops/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBussinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayBussinessDetailLogic {
	return HomestayBussinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessDetailLogic) HomestayBussinessDetail(req types.HomestayBussinessDetailReq) (*types.HomestayBussinessDetailResp, error) {

	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " HomestayBussinessDetail  FindOne db fail ,id  : %d , err : %v", req.Id, err)
	}

	var typeHomestayBusinessBoss types.HomestayBusinessBoss
	if homestayBusiness != nil {

		userResp, err := l.svcCtx.SysRpc.GetUserInfo(l.ctx, &sys.GetUserInfoReq{
			Id: homestayBusiness.UserId,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("get boss info fail"), "get boss info fail ,  userId : %d ,err:%v", homestayBusiness.UserId, err)
		}
		if userResp.User != nil && userResp.User.Id > 0 {
			_ = copier.Copy(&typeHomestayBusinessBoss, userResp.User)
		}
	}

	return &types.HomestayBussinessDetailResp{
		Boss: typeHomestayBusinessBoss,
	}, nil
}
