package result

import (
	"errors"
	"github.com/xu756/qmcy/common/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		var resError *xerr.CodeError
		errors.As(err, &resError)
		httpx.WriteJson(w, 200, Error(resError.ErrCode, resError.ErrMsg))
	}
}
