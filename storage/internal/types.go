package internal

import (
	"context"
	"time"

	model "github.com/TykTechnologies/tyk/storage/internal/model"
	redis6 "github.com/TykTechnologies/tyk/storage/internal/redis6"
	redis7 "github.com/TykTechnologies/tyk/storage/internal/redis7"
)

// RedisDriver implements an abstraction from the underlying client libraries
type RedisDriver interface {
	SetRollingWindow(ctx context.Context, keyName string, per int64, value_override string, pipeline bool) ([]string, error)
	GetRollingWindow(ctx context.Context, keyName string, per int64, pipeline bool) ([]string, error)

	TTL(context.Context, string) (int64, error)
	Get(context.Context, string) (string, error)
	Set(context.Context, string, string, time.Duration) error
	Del(context.Context, string) error

	Keys(context.Context, string) ([]string, error)
	DeleteKeys(ctx context.Context, keys []string) (int64, error)
	DeleteScanMatch(context.Context, string) (int64, error)

	Incr(context.Context, string) (int64, error)
	Decr(context.Context, string) (int64, error)
	MGet(context.Context, []string) (map[string]interface{}, error)
	GetKeysAndValuesWithFilter(context.Context, string) (map[string]interface{}, error)

	Exists(context.Context, string) (int64, error)
	Expire(context.Context, string, time.Duration) error

	FlushAll(context.Context) (bool, error)

	RPush(context.Context, string, string) error

	LRem(context.Context, string, int64, interface{}) (int64, error)
	LRange(context.Context, string, int64, int64) ([]string, error)

	SMembers(context.Context, string) ([]string, error)

	SAdd(context.Context, string, string) error
	SRem(context.Context, string, string) error
	SIsMember(context.Context, string, string) (bool, error)

	ZAdd(context.Context, string, string, float64) (int64, error)
	ZRangeByScoreWithScores(context.Context, string, string, string) (model.ZS, error)
	ZRemRangeByScore(context.Context, string, string, string) (int64, error)

	Publish(context.Context, string, string) (int64, error)
}

// Assert that drivers implement the interface
var _ RedisDriver = &redis6.Driver{}
var _ RedisDriver = &redis7.Driver{}
