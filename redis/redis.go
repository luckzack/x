package redis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"

	"errors"
)

//var RedisConnPool *redis.Pool

type RedisModel struct {
	Addr     string
	Password string
	DB       int
	MaxIdle  int
	Pool     *redis.Pool
}

func New(addr, password string, maxIdle, maxActive int, idleTimeout time.Duration, db int) (*RedisModel, error) {
	if addr == "" {
		return nil, errors.New("Invalid parameters 'addr'.")
	}

	r := RedisModel{
		Addr:     addr,
		Password: password,
		DB:       db,
		MaxIdle:  maxIdle,
	}
	r.Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout, //240 * time.Second,
		Dial: func() (redis.Conn, error) {
			//			c, err := redis.Dial("tcp", redisConfig.Addr)
			c, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(db))

			if err != nil {
				log.Println("[ERROR] Dial redis fail", err)
				return nil, err
			}
			//log.Println("Dial redis succ", addr)
			return c, err
		},
		TestOnBorrow: PingRedis,
	}

	return &r, nil
}

func (r RedisModel) Close() {
	if r.Pool != nil {
		r.Pool.Close()
		r.Pool = nil
	}
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")

	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}

func (r RedisModel) GetConn() redis.Conn {
	return r.Pool.Get()
}

func (r RedisModel) PubSubConn() redis.PubSubConn {
	return redis.PubSubConn{r.Pool.Get()}
}

func (r RedisModel) Publish(subject string, content interface{}) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("PUBLISH", subject, content)
	return
}
