package conn

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "192.168.43.115:6379"
	redisPass = ""
)

// 创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		// 最大连接
		MaxIdle: 50,
		//最大可以同时使用的连接
		MaxActive: 30,
		// 超时时间
		IdleTimeout: 5 * time.Minute,
		Dial: func() (redis.Conn, error) {
			// 打开连接
			conn, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			if len(redisPass) > 0 {
				// 访问认证
				if _, err := conn.Do("AUTH", redisPass); err != nil {
					conn.Close()
					fmt.Println(err.Error())
					return nil, err
				}
			}

			return conn, nil
		},
		// redis健康检查
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
}

func init() {
	pool = newRedisPool()
}
func RedisPool() *redis.Pool {
	return pool
}
