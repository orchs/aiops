package user

import (
	"net/http"

	"aiops/common/result"

	"aiops/app/sys/cmd/api/internal/logic/user"
	"aiops/app/sys/cmd/api/internal/svc"
)

func InfosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewInfosLogic(r.Context(), svcCtx)
		resp, err := l.Infos()
		result.HttpResult(r, w, resp, err)
	}
}
