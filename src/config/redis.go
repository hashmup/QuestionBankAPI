package config

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var RedisTTL = 60 * 60 * 24 * 7

// DBConfig contains the info required for DB connection.
type RedisConfig struct {
	RedisHost string
	RedisPort string
}

// NewConnection uses the env vars to establish db connection and returns it.
func NewRedisConnection(redisConfig RedisConfig) (redis.Conn, error) {
	fmt.Printf("Connecting to :%s:%s", redisConfig.RedisHost, redisConfig.RedisPort)
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.RedisHost, redisConfig.RedisPort))
	if err != nil {
		panic(err)
	}
	fmt.Print("Succeeded to connect to redis.")
	return conn, err
}
