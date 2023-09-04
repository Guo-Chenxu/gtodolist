package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtodolist/app/task/cmd/api/internal/logic"
	"gtodolist/app/task/cmd/api/internal/svc"
	"gtodolist/app/task/cmd/api/internal/types"
)

func task_createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTask_createLogic(r.Context(), svcCtx)
		resp, err := l.Task_create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
