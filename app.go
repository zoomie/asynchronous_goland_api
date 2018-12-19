package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup

type RedisDB struct {
	Addr string
	Password string
	DB int
	conn *redis.Client
}

func (r *RedisDB) MakeConnection() {
	r.conn = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password, // no password set
		DB:       r.DB,  // use default DB
	})
}

func (r *RedisDB) Test() {
	pong, err := r.conn.Ping().Result()
	fmt.Println(pong, err)
}

func (r *RedisDB) SetValue(key string, num string) {
	// defer wg.Done()
	err := r.conn.Set(key, num, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *RedisDB) GetValue(key string) string {
	val, err := r.conn.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func InputValue(db *RedisDB) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.FormValue("key")
		value := r.FormValue("value")
		// wg.Add(1)
		go db.SetValue(key, value)
		fmt.Fprintf(w, "done")
    }
}

func ExtracValue(db *RedisDB) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.FormValue("key")
		value := db.GetValue(key)
		fmt.Fprintf(w, value)
    }
}

func main() {
	database := RedisDB{Addr:"localhost:6379", Password:"", DB:0}
	database.MakeConnection()
	database.Test()
	handleInput := InputValue(&database)
	handleOutput := ExtracValue(&database)
	http.HandleFunc("/SetValue", handleInput)
	http.HandleFunc("/GetValue", handleOutput)
	http.ListenAndServe(":5000", nil)
}
