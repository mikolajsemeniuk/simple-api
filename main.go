package main

import (
	"ex/configuration"
	"ex/handler"
	"ex/router"
	"ex/storage"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	envConfiguration := configuration.EnvConfiguration{}
	if err := envConfiguration.Configure(); err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})

	redisStorage := storage.RedisStorage{
		Client: redisClient,
	}

	articleHandler := &handler.ArticleHandler{
		Redis: redisStorage,
	}

	httpRouter := router.HTTPRouter{
		Configuration:  envConfiguration,
		ArticleHandler: articleHandler,
	}
	if err := httpRouter.Route(); err != nil {
		log.Fatal(err)
	}
}
