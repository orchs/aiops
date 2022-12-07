package homestayOrder

import (
	"net/http"

	"aiops/app/order/cmd/api/internal/logic/homestayOrder"
	"aiops/app/order/cmd/api/internal/svc"
	"aiops/app/order/cmd/api/internal/types"
	"aiops/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserHomestayOrderDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderDetailLogic(r.Context(), ctx)
		resp, err := l.UserHomestayOrderDetail(req)
		result.HttpResult(r, w, resp, err)
	}
}
