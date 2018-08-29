package consistent

import (
	"encoding/json"
	"fmt"
	"sync"
	//"stathat.com/c/consistent"

	consistent "github.com/huichen/consistent_hashing"
)

var redis_nodes = []string{
	"192.168.99.106:7000",
	"192.168.99.157:7000",
	"192.168.99.157:7001",
	"192.168.99.153:7001",
	"192.168.99.153:7002",
	"192.168.99.153:7003",
	"192.168.99.153:7004",
}

var keys = []string{
	"gd", "hn",
}

func main() {
	ring := NewHashRing(redis_nodes, 0)

	m := map[string][]string{}
	for _, node := range redis_nodes {
		m[node] = []string{}
	}

	for _, p := range keys {
		node, _ := ring.GetNode(p)
		//	fmt.Println(p, "->", node)
		m[node] = append(m[node], p)
	}

	fmt.Println(m)

}

type HashRing struct {
	sync.RWMutex
	ring  *consistent.Consistent
	stats map[string]*Stat
}

type Stat struct {
	sync.RWMutex
	TotalHit int
	KeyHit   map[string]int
}

func (s Stat) String() string {
	bytes, _ := json.Marshal(&s)
	return string(bytes)
}

func NewHashRing(nodes []string, numberOfReplicas ...int) *HashRing {
	ret := &HashRing{
		ring:  consistent.New(),
		stats: make(map[string]*Stat),
	}
	if len(numberOfReplicas) > 0 && numberOfReplicas[0] > 0 {
		ret.SetNumberOfReplicas(numberOfReplicas[0])
	}

	ret.SetNodes(nodes)
	return ret
}

func (this *HashRing) SetNodes(nodes []string) {
	for _, node := range nodes {
		this.ring.Add(node)
		this.stats[node] = &Stat{KeyHit: make(map[string]int)}
	}
}

func (this *HashRing) SetNumberOfReplicas(num int) {
	this.ring.NumberOfReplicas = num
}

func (this HashRing) GetNode(pk string) (string, error) {
	node, err := this.ring.Get(pk)

	go this.hit(node, pk)
	return node, err
}

func (this *HashRing) hit(node, key string) {
	this.RLock()
	stat := this.stats[node]
	this.RUnlock()

	stat.Lock()
	stat.TotalHit++
	stat.KeyHit[key]++
	stat.Unlock()
}

func (this HashRing) Stats() map[string]*Stat {
	return this.stats
}
