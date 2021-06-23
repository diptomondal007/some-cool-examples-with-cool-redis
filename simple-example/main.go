package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	var ctx = context.Background()

	redisDb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Run a local redis server. Redis uses 6379 as it's default port.
		DB:   0, // using the db 0. Redis can at most have 15 databases starting from 0.
	})

	fmt.Println(redisDb.Ping(ctx).Result())

	// setting a key value in redis
	stsCmd := redisDb.Set(ctx, "key1", "val1", 0)
	if stsCmd.Err() != nil {
		panic(stsCmd.Err())
	} else {
		log.Println(stsCmd.Result())
	}

	// Getting the value of the key
	val, err := redisDb.Get(ctx, "key1").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(val)

	// Getting the value and trying other fields of *StringCmd
	by, _ := redisDb.Get(ctx, "key1").Bytes()
	log.Println(string(by))

	cmdName := redisDb.Get(ctx, "key1").Name()
	log.Println(cmdName)

	// Getting not existing key
	val, err = redisDb.Get(ctx, "key2").Result()

	if err == redis.Nil {
		log.Println("not found the key")
	} else {
		log.Println(val)
	}


	// SETNX is set a key when the key doesn't exist. Means Set if not exist

	res, err := redisDb.SetNX(ctx, "key", "value", 2*time.Second).Result()
	if err != nil {
		log.Println(err)
	}else {
		log.Println(res)
	}

	// trying to get the key after 2 seconnds.. should get not found
	time.Sleep(2* time.Second)

	valOfTTLSETKEY, err := redisDb.Get(ctx, "key").Result()
	if err == redis.Nil {
		log.Println("key not found")
	}else {
		log.Println(valOfTTLSETKEY)
	}
}
