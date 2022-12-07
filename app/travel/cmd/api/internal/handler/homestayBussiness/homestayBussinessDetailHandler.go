package homestayBussiness

import (
	"net/http"

	"aiops/app/travel/cmd/api/internal/logic/homestayBussiness"
	"aiops/app/travel/cmd/api/internal/svc"
	"aiops/app/travel/cmd/api/internal/types"
	"aiops/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomestayBussinessDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessDetailLogic(r.Context(), ctx)
		resp, err := l.HomestayBussinessDetail(req)
		result.HttpResult(r, w, resp, err)
	}
}
