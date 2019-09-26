package redis

import (
	"fmt"
	"github.com/aon1/slack-horoscope/config"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func New(conf config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		client:client,
	}, nil
}

func (r *Redis) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		fmt.Errorf("Couldnt access redis %v.", err)
	}

	//if err == redis.Nil {
	//	fmt.Errorf("key not found")
	//}

	return val, nil
}

func (r *Redis) Set(key, value string) (string, error) {
	val, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		fmt.Errorf("Couldnt access redis %v.", err)
	}

	//if err == redis.Nil {
	//	fmt.Errorf("key not found")
	//}

	return val, nil
}


