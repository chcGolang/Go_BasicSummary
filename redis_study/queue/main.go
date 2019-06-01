package main

import (
	"Go_BasicSummary/redis_study/conn"
	"github.com/gomodule/redigo/redis"
)

// 生产者
func BatchPushQueue(queueName string, keys []string) (err error) {
	if len(keys) == 0 {
		return
	}
	pool := conn.RedisPool()
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("lpush", redis.Args{}.Add(queueName).AddFlat(keys)...)
	return
}

// 消费者,timeout为0则永久监听
func PopQueue(queueName string, timeout int) (data string, err error) {
	pool := conn.RedisPool()
	con := pool.Get()
	defer con.Close()
	nameAndData, err := redis.Strings(con.Do("brpop", queueName, timeout))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	if len(nameAndData) > 1 {
		data = nameAndData[1]
	}
	return
}
