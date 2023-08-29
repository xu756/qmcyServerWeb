package middleware

import (
	"github.com/xu756/qmcy/app/admin/api/internal/config"
	"github.com/xu756/qmcy/common/result"
	"github.com/xu756/qmcy/common/xctx"
	"github.com/xu756/qmcy/common/xjwt"
	"net/http"
)

type AuthMiddleware struct {
	Jwt *xjwt.Jwt
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Jwt: xjwt.NewJwt(c.Jwt),
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo, err := m.Jwt.GetTokenFromHeader(r, "Authorization")
		if err != nil {
			result.HttpResult(r, w, nil, err)
			return
		}
		ctx := xctx.NewContextForJwt(r.Context(), userInfo)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
