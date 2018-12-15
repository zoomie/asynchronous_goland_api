package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

// Ok so I'm going to put a bit of work into seperating the Redis connection 
// and then passing it into the functions or passing in a pointer.


func test(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// client.Set("test", "test1")
	// val, _ := client.Get("test").Result()
	fmt.Println(client)
	// r.ParseForm()
	// fmt.Println(r.FormValue("key"))
	// fmt.Println(r.FormValue("value"))
	// fmt.Println(r.Form)
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
	http.HandleFunc("/test", test)
	http.HandleFunc("/setValue", setValue)
	http.HandleFunc("/getValue", getValue)
	http.ListenAndServe(":5000", nil)
}
