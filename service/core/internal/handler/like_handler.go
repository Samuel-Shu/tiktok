package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-tiktok/service/core/internal/logic"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"
)

func LikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		userId, _ := strconv.Atoi(r.Header.Get("UserId"))
		req.UserId = uint(userId)

		l := logic.NewLikeLogic(r.Context(), svcCtx)
		resp, err := l.Like(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
