// package main

// import (
// 	"fmt"
// 	"sync"

// 	"github.com/go-redis/redis"
// )

// var wg sync.WaitGroup

// type RedisDB struct {
// 	Addr string
// 	Password string
// 	DB int
// 	conn *redis.Client
// }

// func (r *RedisDB) MakeConnection() {
// 	r.conn = redis.NewClient(&redis.Options{
// 		Addr:     r.Addr,
// 		Password: r.Password, // no password set
// 		DB:       r.DB,  // use default DB
// 	})
// }

// func (r *RedisDB) Test() {
// 	defer wg.Done()
// 	pong, err := r.conn.Ping().Result()
// 	fmt.Println(pong, err)
// }

// func (r *RedisDB) SetValue(key string, num string) {
// 	err := r.conn.Set(key, num, 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (r *RedisDB) GetValue(key string) string {
// 	val, err := r.conn.Get(key).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return val
// }

// func p(words string) {
// 	defer wg.Done()
// 	fmt.Println(words)
// }

// func main() {
// 	database := RedisDB{Addr:"localhost:6379", Password:"", DB:0}
// 	database.MakeConnection()
// 	for i:=0;i<100000;i++ {
// 		wg.Add(1)
// 		go database.Test()
// 	}

// 	wg.Wait()
// 	fmt.Println("done")
// }