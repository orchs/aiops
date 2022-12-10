package user

import (
	"net/http"

	"aiops/common/result"

	"aiops/app/sys/cmd/api/internal/logic/user"
	"aiops/app/sys/cmd/api/internal/svc"
	"aiops/app/sys/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		result.HttpResult(r, w, resp, err)
	}
}
