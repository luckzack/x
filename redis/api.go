package redis

//通用接口，不涉及业务

func (r RedisModel) Keys(keyword string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("keys", keyword+"*")
	return
}

func (r RedisModel) Set(key string, value interface{}) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SET", key, value)
	return
}

func (r RedisModel) Expire(key string, seconds uint) {
	rc := r.Pool.Get()
	defer rc.Close()
	rc.Do("EXPIRE", key, seconds)
}

func (r RedisModel) Append(key, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("APPEND", key, value)
	return
}

func (r RedisModel) Get(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("GET", key)
	return
}

func (r RedisModel) Decr(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("DECR", key)
	return
}

func (r RedisModel) DecrBy(key string, num int) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("DECRBY", key, num)
	return
}

func (r RedisModel) Incr(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("INCR", key)
	return
}

func (r RedisModel) IncreBy(key string, num int) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("INCRBY", key, num)
	return
}

func (r RedisModel) Del(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("DEL", key)
	return
}

func (r RedisModel) Clear(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	reply, err = rc.Do("SET", key, "")
	return
}

//Hash

func (r RedisModel) HSet(key, field string, value interface{}) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HSET", key, field, value)
	return
}

func (r RedisModel) HGet(key, field string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HGET", key, field)
	return
}

//func (r RedisModel)  HMGet(key string, fields []uint ) (reply interface{}, err error){
//	rc := r.Pool.Get()
//	defer rc.Close()
//	//arr := []int{35, 36}
//	reply, err = rc.Do("HMGET", key, fields...)
//	return
//}

func (r RedisModel) HMSet(key string, m map[string]map[string][][]string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()
	for k, v := range m {
		bytes, _ := json.Marshal(&v)
		rc.Do("HMSET", key, k, bytes)
	}

	return
}

// clear：取出数据后是否清空hash
func (r RedisModel) HGetAllValues(key string, clear bool) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HVALS", key)
	if clear {
		rc.Do("DEL", key)
	}
	return
}
func (r RedisModel) HGetAll(key string, clear bool) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HGETALL", key)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

func (r RedisModel) HLen(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HLEN", key)
	return
}

func (r RedisModel) HDel(key, field string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HDEL", key, field)
	return
}

func (r RedisModel) HIncr(key, field string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HINCRBY", key, field, 1)
	return
}

func (r RedisModel) HDecr(key, field string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("HINCRBY", key, field, -1)
	return
}

//List

func (r RedisModel) LPush(key, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LPUSH", key, value)
	return
}

// clear：取出数据后是否清空list
func (r RedisModel) LGetAll(key string, clear bool) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LRANGE", key, 0, -1)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

// 移除list中所有与value相等的值
func (r RedisModel) LREM(key string, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LREM", key, 0, value)

	return
}

func (r RedisModel) LLen(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LLEN", key)

	return
}

func (r RedisModel) LFIRST(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LINDEX", key, 0)

	return
}

// 移除并返回列表 key 的尾元素。
func (r RedisModel) RPOP(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("RPOP", key)

	return
}

// 将一个或多个值 value 插入到列表 key 的表尾(最右边)。
func (r RedisModel) RPUSH(key string, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("RPUSH", key, value)

	return
}

//移除并返回列表 key 的头元素
func (r RedisModel) LPOP(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("LPOP", key)

	return
}

//移除并返回列表 key 的头元素
func (r RedisModel) LPOPN(key string, n int) (replies []interface{}) {
	rc := r.Pool.Get()
	defer rc.Close()

	for i := 0; i < n; i++ {
		reply, err := rc.Do("LPOP", key)
		if err == nil {
			replies = append(replies, reply)
		}
	}

	return
}

/****
* SET 无序集合
 */

// 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
func (r RedisModel) SADD(key, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SADD", key, value)

	return
}

func (r RedisModel) SADD_KEYS(keys []string, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	for _, key := range keys {
		reply, err = rc.Do("SADD", key, value)
	}

	return
}

func (r RedisModel) SADD_VALUES(key string, values ...string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	for _, value := range values {
		reply, err = rc.Do("SADD", key, value)
	}

	return
}

//移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略。
func (r RedisModel) SREM(key, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SREM", key, value)
	return
}

func (r RedisModel) SREM_KEYS(keys []string, value string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	for _, key := range keys {
		reply, err = rc.Do("SREM", key, value)
	}

	return
}

func (r RedisModel) SREM_VALUES(key string, values ...string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	for _, value := range values {
		reply, err = rc.Do("SREM", key, value)
	}

	return
}

// 返回集合 key 中的所有成员。
func (r RedisModel) SMEMBERS(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SMEMBERS", key)
	return
}

// 返回集合 key 的基数(集合中元素的数量)。
func (r RedisModel) SCARD(key string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SCARD", key)
	return
}

// 判断 member 元素是否集合 key 的成员。
func (r RedisModel) SISMEMBER(key, member string) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("SISMEMBER", key)
	return
}
