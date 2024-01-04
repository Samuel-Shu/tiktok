package handler

import (
	"errors"
	"log"
	"mini-tiktok/core/helper"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-tiktok/core/internal/logic"
	"mini-tiktok/core/internal/svc"
	"mini-tiktok/core/internal/types"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		_, _, err := r.FormFile("data")
		if err != nil {
			log.Println(err)
			httpx.ErrorCtx(r.Context(), w, errors.New("请上传文件"))
			return
		}

		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}

		req.PlayURL = cosPath

		l := logic.NewPublishLogic(r.Context(), svcCtx)
		userId, _ := strconv.Atoi(r.Header.Get("UserId"))

		resp, err := l.Publish(&req, int64(userId))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
