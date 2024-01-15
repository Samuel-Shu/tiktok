package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-tiktok/service/core/internal/logic"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"
)

func GetFansHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFansRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFansLogic(r.Context(), svcCtx)
		resp, err := l.GetFans(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
