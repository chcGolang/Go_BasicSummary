package pub_sub

import (
	"Go_BasicSummary/redis_study/conn"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 发布者
func Publish(channleName string, message string) (subscribeNum int, err error) {
	pool := conn.RedisPool().Get()
	defer pool.Close()
	subscribeNum, err = redis.Int(pool.Do("publish", channleName, message))
	return
}

// 订阅者
func Subscribe(channleName ...interface{}) (err error) {
	pool := conn.RedisPool().Get()
	defer pool.Close()
	subConn := redis.PubSubConn{pool}
	if err = subConn.Subscribe(channleName...); err != nil {
		return
	}
	for {
		switch v := subConn.Receive().(type) {
		case redis.Message: // 订阅的消息
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription: // 订阅的配置信息
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		}
	}
}
