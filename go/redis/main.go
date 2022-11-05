package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	// result 方式
	b, err := rdb.HMSet(context.Background(), "a", "m1", 1, "m2", 2).Result()

	log.Printf("%v\t%v\n", b, err)

	// pipe 执行多条命令和获取结果
	pipe := rdb.Pipeline()
	pipe.Set(ctx, "pipe1", "1", time.Hour)
	pipe.Set(ctx, "pipe2", "2", time.Hour)
	pipe.Set(ctx, "pipe3", "3", time.Hour)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, cmd := range cmds {
		v, err := cmd.(*redis.StatusCmd).Result()
		log.Printf("%v\t%v\n", v, err)
	}

	// 结构体
	type Person struct {
		Name string `redis:"name"`
		Age  int    `redis:"age"`
	}

	rdb.HMSet(ctx, "p1", "name", "leo", "age", 18)

	var p Person
	rdb.HGetAll(ctx, "p1").Scan(&p)
	log.Printf("%T\t%v\n", p, p)
	// main.Person {leo 18}

	// pipe + 结构体
	rdb.HMSet(ctx, "p1", "name", "leo", "age", 18)
	rdb.HMSet(ctx, "p2", "name", "leo2", "age", 18)
	rdb.HMSet(ctx, "p3", "name", "leo3", "age", 18)

	pipe = rdb.Pipeline()
	pipe.HGetAll(ctx, "p1")
	pipe.HGetAll(ctx, "p2")
	pipe.HGetAll(ctx, "p3")
	cmds, err = pipe.Exec(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, cmd := range cmds {
		var p Person
		err := cmd.(*redis.StringStringMapCmd).Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%T\t%v\n", p, p)
	}
}
