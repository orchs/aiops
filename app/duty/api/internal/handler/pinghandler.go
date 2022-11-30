package handler

import (
	"aiops/app/duty/api/internal/logic"
	"aiops/app/duty/api/internal/svc"
	"aiops/common/response"
	"net/http"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		response.Response(w, nil, err)

	}
}
