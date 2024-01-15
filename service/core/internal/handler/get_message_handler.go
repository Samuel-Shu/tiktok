package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-tiktok/service/core/internal/logic"
	"mini-tiktok/service/core/internal/svc"
	"mini-tiktok/service/core/internal/types"
)

func GetMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		get := r.Header.Get("UserId")
		userId, _ := strconv.Atoi(get)
		req.UserId = uint(userId)

		l := logic.NewGetMessageLogic(r.Context(), svcCtx)
		resp, err := l.GetMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
