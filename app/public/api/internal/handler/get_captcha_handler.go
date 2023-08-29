package handler

import (
	"github.com/xu756/qmcy/common/result"
	"net/http"

	"github.com/xu756/qmcy/app/public/api/internal/logic"
	"github.com/xu756/qmcy/app/public/api/internal/svc"
	"github.com/xu756/qmcy/app/public/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
		result.HttpResult(r, w, resp, err)
	}
}
