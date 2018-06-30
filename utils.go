package utils

import (
	"math/rand"
	"time"
)

// 最大值，最小值，单位ms
func SleepRandomDuration(max int64, least int64) {
	ns := max * 1000000
	// 以当前时间为随机数种子，如果所有ops-updater在同一时间启动，系统时间是相同的，那么随机种子就是一样的
	// 问题不大，批量ssh去启动ops-updater的话也是一个顺次的过程
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := time.Duration(r.Int63n(ns)+least*1000000) * time.Nanosecond
	time.Sleep(d)
}
