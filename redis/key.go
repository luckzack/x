package redis

//与业务相关的redis操作features

import (
	"athena-v4/common/model"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"fmt"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (

	// 所有任务的概览，只做简单查询
	Key_preload_alltasks = "alltasks" //hash,field=task_id

	// 所有任务的详情，客户端来从这里取
	Key_preload_tasks_cid_system = "tasks:cid=%v:platform=%v" //key=tasks:cid:platform,field=task_id
	Key_preload_taskkeys         = "taskkeys"                 // taskid存在哪些key中，当更新cache的时候，要把旧的删掉

	// 队列中，sdk请求来之后如果分配的任务id，那么这个任务分配记录也要记录到mysql（记录taskid-uid，finish_cnt）
	Key_preload_queue_task = "queue_task" //field:task_id ,value:jsontask

	// 队列中，sdk请求记录，要取出来存到echos表
	Key_preload_queue_request  = "queue_request"    //[]request
	Key_queue_request_platform = "queue_request_%v" //[]request

	// 记录每个任务发给了那些uid, LIST
	Key_preload_task_uids = "task_uids:task_id=%v"

	// 记录任务的完成数、计划完成数、上一次被使用的时间，hash
	Key_preload_tasks_summary = "tasks_summary"

	// 终端cid到客户cid的关系
	Key_subscriptions = "subscriptions"

	// CP客户信息（cid，domain等）
	Key_cps = "cps"

	// 预加载终端信息
	Key_terminals = "terminals"

	// 预加载任务分配模式
	Key_allocateModes = "allocateModes"

	// 预加载策略
	Key_rules = "rules"

	// 预加载uid最新记录
	Key_uidRecords = "uidRecords"

	// domain配置
	Key_domains = "domains"

	//计划作业
	Key_jobs = "jobs"

	// url参数
	Key_urlArgs = "url_args"
)

// sdk任务的队列操作
func (r RedisModel) AppendTaskIntoQueue(jtask *model.QueueTask) error {

	bytes, _ := json.Marshal(jtask)

	_, err := r.RPUSH(Key_preload_queue_task, string(bytes))
	return err
}

func (r RedisModel) ShiftTaskFromQueue() (*model.QueueTask, error) {
	reply, err := r.LPOP(Key_preload_queue_task)
	//log.Println("ShiftTaskFromQueue --->", reply, err)
	if err != nil {
		return nil, err
	}

	if reply == nil {
		return nil, nil
	}

	if v, ok := reply.([]byte); ok {
		jtask := model.QueueTask{}
		err := json.Unmarshal(v, &jtask)

		return &jtask, err
	}

	return nil, errors.New("reply to bytes err")
}

func (r RedisModel) ShiftsTaskFromQueue(n int) []*model.QueueTask {
	tasks := []*model.QueueTask{}

	replies := r.LPOPN(Key_preload_queue_task, n)

	cnt := len(replies)
	for i := 0; i < cnt; i++ {

		if v, ok := replies[i].([]byte); ok {
			qtask := model.QueueTask{}
			err := json.Unmarshal(v, &qtask)

			if err == nil {
				tasks = append(tasks, &qtask)
			}
		}
	}

	return tasks
}

// sdk请求的队列操作
func (r RedisModel) AppendRequestIntoQueue(req *model.QueueRequest) error {
	bytes, _ := json.Marshal(req)
	//	_, err := RPUSH(Key_preload_queue_request, string(bytes))
	_, err := r.RPUSH(fmt.Sprintf(Key_queue_request_platform, req.Platform), string(bytes))
	return err
}

func (r RedisModel) ShiftOneRequestFromQueue() (*model.QueueRequest, error) {
	reply, err := r.LPOP(Key_preload_queue_request)
	if err != nil {
		return nil, err
	}

	if reply == nil {
		return nil, nil
	}

	if v, ok := reply.([]byte); ok {
		qreq := model.QueueRequest{}
		err := json.Unmarshal(v, &qreq)

		return &qreq, err
	}
	return nil, errors.New("reply to bytes err")
}

func (r RedisModel) ShiftRequestsFromQueue(n int) []*model.QueueRequest {
	//log.Println("ShiftRequestsFromQueue-->", n)
	qreqs := []*model.QueueRequest{}

	replies := r.LPOPN(Key_preload_queue_request, n)

	cnt := len(replies)
	for i := 0; i < cnt; i++ {

		if v, ok := replies[i].([]byte); ok {
			qreq := model.QueueRequest{}
			err := json.Unmarshal(v, &qreq)

			if err == nil {
				qreqs = append(qreqs, &qreq)
			}
		}
	}

	return qreqs
}

func (r RedisModel) ShiftRequestsFromQueueByPlatform(n, platform int) []*model.QueueRequest {
	//log.Println("ShiftRequestsFromQueue-->", n)
	qreqs := []*model.QueueRequest{}

	replies := r.LPOPN(fmt.Sprintf(Key_queue_request_platform, platform), n)

	cnt := len(replies)
	for i := 0; i < cnt; i++ {

		if v, ok := replies[i].([]byte); ok {
			qreq := model.QueueRequest{}
			err := json.Unmarshal(v, &qreq)

			if err == nil {
				qreqs = append(qreqs, &qreq)
			}
		}
	}

	return qreqs
}

//
func (r RedisModel) SetTask(field, value string) error {
	_, err := r.HSet(Key_preload_alltasks, field, value)
	return err
}

func (r RedisModel) SetTaskWithKey(key, field, value string) error {
	_, err := r.HSet(key, field, value)

	return err
}

func (r RedisModel) GetAllTasks() ([]*model.JsonTask, error) {
	reply, err := r.HGetAllValues(Key_preload_alltasks, false)
	if err != nil {
		return nil, err
	}

	tasks := []*model.JsonTask{}

	if v, ok := reply.([]interface{}); ok {

		for _, vv := range v {
			if v2, ok2 := vv.([]byte); ok2 {

				task := model.JsonTask{}
				err := json.Unmarshal(v2, &task)

				if err == nil {
					tasks = append(tasks, &task)
				}
			}
		}
	}

	return tasks, nil
}

func (r RedisModel) GetAllTasksWithKey(key string) ([]*model.JsonTask, error) {
	reply, err := r.HGetAllValues(key, false)
	if err != nil {
		return nil, err
	}

	tasks := []*model.JsonTask{}

	if v, ok := reply.([]interface{}); ok {

		for _, vv := range v {
			if v2, ok2 := vv.([]byte); ok2 {

				task := model.JsonTask{}
				err := json.Unmarshal(v2, &task)

				if err == nil {
					tasks = append(tasks, &task)
				}
			}
		}
	}

	return tasks, nil
}

func (r RedisModel) GetTasksByIds(ids []uint) (tasks []*model.JsonTask) {

	for _, id := range ids {
		id_str := strconv.Itoa(int(id))
		reply, err := r.HGet(Key_preload_alltasks, id_str)
		if err == nil {
			continue
		}

		if v, ok := reply.([]byte); ok {
			task := model.JsonTask{}
			err := json.Unmarshal(v, &task)

			if err == nil {
				tasks = append(tasks, &task)
			}
		}

	}
	return
}

func (r RedisModel) GetTasksByIdThenDelete(ids []uint) (tasks []*model.JsonTask) {

	for _, id := range ids {
		id_str := strconv.Itoa(int(id))
		reply, err := r.HGet(Key_preload_alltasks, id_str)

		if err != nil {
			continue
		}

		if v, ok := reply.([]byte); ok {
			task := model.JsonTask{}
			err := json.Unmarshal(v, &task)

			if err == nil {
				tasks = append(tasks, &task)
			}
		}

		r.HDel(Key_preload_alltasks, id_str)

	}
	return
}

func (r RedisModel) DeleteWithKey(key string, id uint) {
	id_str := strconv.Itoa(int(id))
	r.HDel(key, id_str)
}

///
func (r RedisModel) SetTaskKeys(task_id uint, keys []string) {
	id_str := strconv.Itoa(int(task_id))
	bytes, _ := json.Marshal(&keys)
	r.HSet(Key_preload_taskkeys, id_str, string(bytes))
}

func (r RedisModel) GetTaskKeys(task_id uint) ([]string, error) {
	id_str := strconv.Itoa(int(task_id))
	reply, err := r.HGet(Key_preload_taskkeys, id_str)
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return nil, nil
	}

	if v, ok := reply.([]byte); ok {
		keys := []string{}
		err = json.Unmarshal(v, &keys)
		return keys, err
	}

	err = errors.New("reply to bytes err")
	return nil, err
}

///// 从redis.preload_tasking中取出taskid发给了哪些uid
//func (r RedisModel)GetTasking(task_id int) (string, error) {
//	id_str := strconv.Itoa(task_id)
//	reply, err := HGet(Key_preload_tasking, id_str)
//	if err != nil {
//		return "", err
//	}
//	if reply == nil {
//		return "", nil
//	}
//	if v, ok := reply.([]byte); ok {
//		return string(v), err
//	}
//	err = errors.New("reply to bytes err")
//	return "", err
//}

//
func (r RedisModel) GetTaskUids(task_id uint) (uids []string, err error) {
	reply, err := r.LGetAll(fmt.Sprintf(Key_preload_task_uids, task_id), false)
	if err != nil {
		return nil, err
	}

	if v1, ok := reply.([]interface{}); ok {
		for _, v11 := range v1 {
			if bytes, ok := v11.([]uint8); ok {
				uids = append(uids, string(bytes))
			}
		}

	}

	return
}

func (r RedisModel) AddTaskUid(task_id uint, uid string) error {
	_, err := r.RPUSH(fmt.Sprintf(Key_preload_task_uids, task_id), uid)
	return err
}

func (r RedisModel) DelTaskUids(task_id uint) error {
	_, err := r.Del(fmt.Sprintf(Key_preload_task_uids, task_id))
	return err
}
