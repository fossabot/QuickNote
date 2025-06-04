package rps

import (
	"context"
	"fmt"
	"time"

	"github.com/Sn0wo2/QuickNote/rds"
	"github.com/go-redis/redis/v8"
)

var Instance = NewRPS()

type RPS struct {
	ctx context.Context
}

func NewRPS() *RPS {
	return &RPS{
		ctx: context.Background(),
	}
}

func (r *RPS) RecordRequest() error {
	now := time.Now().Unix()

	pipe := rds.Instance.Get().Pipeline()

	pipe.ZAdd(r.ctx, GetKey(), &redis.Z{
		Score:  float64(now),
		Member: fmt.Sprintf("%d-%d", now, time.Now().UnixNano()),
	})

	pipe.ZRemRangeByScore(r.ctx, GetKey(), "-inf", fmt.Sprintf("%d", now-60))

	pipe.Expire(r.ctx, GetKey(), 2*time.Minute)

	_, err := pipe.Exec(r.ctx)
	return err
}

func (r *RPS) GetRequestsPerMinute() (int64, error) {
	now := time.Now().Unix()

	count, err := rds.Instance.Get().ZCount(r.ctx, GetKey(),
		fmt.Sprintf("%d", now-60),
		fmt.Sprintf("%d", now)).Result()

	return count, err
}

func (r *RPS) GetRequestsPerSecond() (float64, error) {
	count, err := r.GetRequestsPerMinute()
	if err != nil {
		return 0, err
	}

	return float64(count) / 60.0, nil
}

func GetKey() string {
	return "rps:requests"
}
