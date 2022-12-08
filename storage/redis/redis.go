package redis

import (
	// "context"
	"api-exam/storage/repo"

	"github.com/gomodule/redigo/redis"
	// "github.com/go-redis/redis/v8"
	// "time"
)

type RedisRepo struct {
	rdb *redis.Pool
}

func NewRedisRepo(rdb *redis.Pool) repo.InMemorystorageI {
	return &RedisRepo{
		rdb: rdb,
	}
}

func (r *RedisRepo) Exists(key string) (interface{}, error) {
	conn := r.rdb.Get()
	defer conn.Close()
	return conn.Do("EXISTS", key)
}

func (r *RedisRepo) SetWithTTL(key, value string, seconds int) (err error) {
	conn := r.rdb.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, seconds, value)
	return
}

func (r *RedisRepo) Get(key string) (interface{}, error) {
	conn := r.rdb.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}
