package homestayBussiness

import (
	"context"
	"github.com/Masterminds/squirrel"

	"aiops/app/travel/cmd/api/internal/svc"
	"aiops/app/travel/cmd/api/internal/types"
	"aiops/app/travel/model"
	"aiops/app/sys/cmd/rpc/sys"
	"aiops/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodBossLogic {
	return GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(req types.GoodBossReq) (*types.GoodBossResp, error) {

	whereBuilder := l.svcCtx.HomestayActivityModel.RowBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityGoodBusiType,
		"row_status": model.HomestayActivityUpStatus,
	})
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, 0, 10, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get GoodBoss db err. rowType: %s ,err : %v", model.HomestayActivityGoodBusiType, err)
	}

	var resp []types.HomestayBusinessBoss
	if len(homestayActivityList) > 0 {

		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		}, func(item interface{}, writer mr.Writer, cancel func(error)) {
			id := item.(int64)

			userResp, err := l.svcCtx.SysRpc.GetUserInfo(l.ctx, &sys.GetUserInfoReq{
				Id: id,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("GoodListLogic GoodList fail userId : %d ,err:%v", id, err)
				return
			}
			if userResp.User != nil && userResp.User.Id > 0 {
				writer.Write(userResp.User)
			}
		}, func(pipe <-chan interface{}, cancel func(error)) {

			for item := range pipe {
				var typesHomestayBusiness types.HomestayBusinessBoss
				_ = copier.Copy(&typesHomestayBusiness, item)

				// compute star todo
				resp = append(resp, typesHomestayBusiness)
			}
		})
	}

	return &types.GoodBossResp{
		List: resp,
	}, nil
}
