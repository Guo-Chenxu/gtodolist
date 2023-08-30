package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtodolist/app/user/cmd/api/internal/logic"
	"gtodolist/app/user/cmd/api/internal/svc"
	"gtodolist/app/user/cmd/api/internal/types"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
