package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("key"))
	fmt.Println(r.FormValue("value"))
	fmt.Println(r.Form)
}

func setValue(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r.ParseForm()
	key := r.FormValue("key")
	value := r.FormValue("value")
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("set value done")
}

func getValue(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r.ParseForm()
	key := r.FormValue("key")
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Get value done", val)
}

func main() {
	// http.HandleFunc("/test", test)
	http.HandleFunc("/setValue", setValue)
	http.HandleFunc("/getValue", getValue)
	http.ListenAndServe(":5000", nil)
}
