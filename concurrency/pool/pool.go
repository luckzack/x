package pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DefaultPoolSize = 100
	default_timeout = time.Second
)

type Pool struct {
	Name    string
	Size    uint
	TTL     time.Duration //second
	Timeout time.Duration //second
	pool    chan bool
}

type Connection struct {
	ctx      *Pool
	released bool
}

func (p Pool) Map() map[string]interface{} {
	m := map[string]interface{}{
		//"name": p.Name,
		"use":     p.Use(),
		"idle":    p.Idle(),
		"timeout": p.Timeout.String(),
		"ttl":     p.TTL.String(),
	}

	return m
}

func (p Pool) String() string {
	m := p.Map()
	bytes, _ := json.Marshal(&m)
	return string(bytes)
}

func New(name string, size uint, timeout uint, ttl ...int) *Pool {

	p := Pool{
		Name: name,
		Size: size,
		//TTL:     time.Second * time.Duration(ttl),
		//Timeout: time.Second * time.Duration(timeout),
		pool: make(chan bool, size),
	}
	if len(ttl) > 0 && ttl[0] > 0 {
		p.TTL = time.Second * time.Duration(ttl[0])
	}

	if timeout > 0 {
		p.Timeout = time.Second * time.Duration(timeout)
	} else {
		p.Timeout = default_timeout
	}

	return &p
}

func (p Pool) Use() int {
	return len(p.pool)
}

func (p Pool) Idle() int {
	return int(p.Size) - len(p.pool)
}

func (p Pool) Print() string {
	return fmt.Sprintf("Pool{name: %s, use: %d, idle: %d}", p.Name, p.Use(), p.Idle())
}

func (p *Pool) GetConn() (*Connection, error) {

	timeout := time.NewTimer(p.Timeout)

	for {
		select {
		case p.pool <- true:

			connection := &Connection{ctx: p}
			if p.TTL.Seconds() > 0 {
				go func(conn *Connection) {
					time.Sleep(p.TTL)
					conn.Release()
				}(connection)
			}

			return connection, nil
		case <-timeout.C:
			return nil, errors.New(fmt.Sprintf("try again later. %s", p.Print()))
		}
	}

}

func (c *Connection) Release() {

	if !c.released {
		<-c.ctx.pool
	}
	c.released = true
	//fmt.Println("conn.released ->idle:", c.ctx.Idle())
}
