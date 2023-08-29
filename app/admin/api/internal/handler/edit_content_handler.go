package handler

import (
	"net/http"

	"github.com/xu756/qmcy/app/admin/api/internal/logic"
	"github.com/xu756/qmcy/app/admin/api/internal/svc"
	"github.com/xu756/qmcy/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Content
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEditContentLogic(r.Context(), svcCtx)
		resp, err := l.EditContent(&req)
		result.HttpResult(r, w, resp, err)
	}
}
