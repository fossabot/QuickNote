package rds

import (
	"github.com/go-redis/redis/v8"
)

var Instance = New()

type Redis struct {
	client *redis.Client
}

func New() *Redis {
	return &Redis{}
}

func (r *Redis) Init(options *redis.Options) {
	r.client = redis.NewClient(options)
}

func (r *Redis) Get() *redis.Client {
	return r.client
}

func (r *Redis) Close() error {
	return r.client.Close()
}
