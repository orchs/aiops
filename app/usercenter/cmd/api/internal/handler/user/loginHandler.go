package user

import (
	"net/http"

	"aiops/app/usercenter/cmd/api/internal/logic/user"
	"aiops/app/usercenter/cmd/api/internal/svc"
	"aiops/app/usercenter/cmd/api/internal/types"
	"aiops/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		//
		l := user.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		result.HttpResult(r, w, resp, err)
	}
}