package storage

import (
	"github.com/go-redis/redis"
)

type RedisDocument interface {
	String() string
}

type RedisStorage struct {
	Client *redis.Client
}

func (s *RedisStorage) Read(key string) ([]byte, error) {
	result, err := s.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(result), err
}

func (s *RedisStorage) Write(key string, document RedisDocument) error {
	return s.Client.Set(key, document.String(), 0).Err()
}
