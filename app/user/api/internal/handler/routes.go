// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/xu756/qmcy/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/getUserInfo",
					Handler: GetUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/app/user"),
	)
}
