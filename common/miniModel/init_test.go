package miniModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var DbSource = "postgres://root:fyU0bj2KLndYN6CKYubb@localhost:15432/qmcy_mini?sslmode=disable"
var RedisConf = []cache.NodeConf{
	{
		RedisConf: redis.RedisConf{
			Host: "localhost:16379",
			Pass: "Y5dg5tg8050oigInC30sf",
			Type: "node",
		},
		Weight: 100,
	},
}
