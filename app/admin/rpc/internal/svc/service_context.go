package svc

import (
	"github.com/xu756/qmcy/app/admin/rpc/internal/config"
	"github.com/xu756/qmcy/common/miniModel"
	"github.com/xu756/qmcy/common/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config           config.Config
	RedisClient      *redis.Redis
	UserModel        model.UserModel
	GroupModel       model.GroupModel
	UserRoleModel    model.UserRoleModel
	UserGroupModel   model.UserGroupModel
	MiniContentModel miniModel.ContentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	redisClient, err := redis.NewRedis(c.RedisConf)
	conn := postgres.New(c.DbSource)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		RedisClient:      redisClient,
		UserModel:        model.NewUserModel(conn, c.Cache),
		GroupModel:       model.NewGroupModel(conn, c.Cache),
		UserRoleModel:    model.NewUserRoleModel(conn, c.Cache),
		UserGroupModel:   model.NewUserGroupModel(conn, c.Cache),
		MiniContentModel: miniModel.NewContentsModel(postgres.New(c.MiniDbSource), c.Cache),
	}
}
