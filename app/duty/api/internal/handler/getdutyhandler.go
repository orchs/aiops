package handler

import (
	"aiops/app/duty/api/internal/logic"
	"aiops/app/duty/api/internal/svc"
	"aiops/app/duty/api/internal/types"
	"aiops/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func getDutyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DutyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetDutyLogic(r.Context(), svcCtx)
		resp, err := l.GetDuty(&req)
		response.Response(w, resp, err)

	}
}
