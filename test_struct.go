package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Setup struct {
	conn *redis.Client
}

func (s Setup) test() {
	pong, err := s.conn.Ping().Result()
	fmt.Println(pong, err)
	fmt.Println("success")
}

func (s Setup) SetValue(key, value string) {
	err := s.conn.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("the key was set", key)
}

func (s Setup) GetValue(key string) string {
	val, err := s.conn.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redis_db := Setup{client}
	redis_db.SetValue("key5", "value100")
	val := redis_db.GetValue("key5")
	fmt.Println("the returned key is", val)
}

// func redis_func() *redis.Client {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// DB:       0,  // use default DB
// 	})
// 	return client
// }
