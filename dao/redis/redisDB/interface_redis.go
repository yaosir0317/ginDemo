package redisDB

import (
	"ginDemo/dao/redis/connections"

	"github.com/go-redis/redis/v7"
)

type Interface1 interface {
	GetKey() (string, error)
	MGetKey() ([]string, error)
	GetRemovePreQ(d string) (*[]string, error)
	HGetInfo(articleId string) error
}

type impl struct {
	client *redis.Client
}

type Interface interface {
	Interface1
}

var _ Interface = (*impl)(nil)

func NewRedisInterface() *impl {
	return &impl{
		client: connections.NewRedisConnection(),
	}
}
