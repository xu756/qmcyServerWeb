package svc

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/xu756/qmcy/app/public/api/internal/config"
	"github.com/xu756/qmcy/app/public/rpc/public"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"net/url"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	PublicRpc   public.Public
	CosClient   *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	u, _ := url.Parse(c.Cos.Url)
	b := &cos.BaseURL{BucketURL: u}
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		PublicRpc:   public.NewPublic(zrpc.MustNewClient(c.Public)),
		CosClient: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  c.Cos.SecretID,
				SecretKey: c.Cos.SecretKey,
			},
		}),
	}
}
