package config

import (
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

var (
	once_re sync.Once
	p_re    *redis.Pool
)

func Conn_Re_DB() *redis.Pool {
	once_re.Do(func() {
		p_re = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "redis:6379")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}
				return conn, err
			},
		}
	})

	return p_re
}
