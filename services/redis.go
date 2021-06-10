package services

import (
	"encoding/json"
	"time"

	"github.com/charlesonunze/busha-test/model"
	"github.com/go-redis/redis"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
)

func StoreInCache(data []model.Movie, key string) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = client.Set(key, json, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func FetchFromCache(key string) ([]model.Movie, error) {
	results := []model.Movie{}
	val, err := client.Get(key).Result()

	switch {
	case err == redis.Nil:
		return results, nil
	case err != nil:
		return results, err
	case val == "":
		return results, err
	}

	err = json.Unmarshal([]byte(val), &results)
	if err != nil {
		return results, err
	}

	return results, nil
}