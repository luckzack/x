package redis

import (
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	r, err := New("192.168.1.211:6379", "99101691", 10, 100, time.Minute, 4)

	if err != nil {
		log.Panic(err.Error())
	}

	conn := r.GetConn()

	defer conn.Close()

	vals := []int{1, 2, 3}
	reply, err := conn.Do("RPUSH", "key", vals)
	log.Println(reply, err)

}
