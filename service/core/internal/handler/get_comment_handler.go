package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-tiktok/service/core/internal/logic"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"
	"net/http"
)

func GetCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentRequest

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCommentLogic(r.Context(), svcCtx)
		resp, err := l.GetComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
