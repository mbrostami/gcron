package db

import (
	"encoding/json"
	"sort"
	"time"
	"unsafe"

	pb "github.com/mbrostami/gcron/internal/grpc"
	"github.com/rs/xid"
	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
	log "github.com/sirupsen/logrus"
)

// LedisDB database
type LedisDB struct {
	ledis *ledis.Ledis
	db    *ledis.DB
}

// NewLedis create ledisdb instance
func NewLedis() *LedisDB {
	cfg := config.NewConfigDefault()
	ledis, err := ledis.Open(cfg)
	if err != nil {
		log.Fatalf("DB Connect error! %v", err)
	}
	db, _ := ledis.Select(0)
	return &LedisDB{db: db, ledis: ledis}
}

// Store data in db
func (l LedisDB) Store(uid uint32, task *pb.Task) (string, error) {
	byteKeys := (*[4]byte)(unsafe.Pointer(&uid))[:] // 32 bit id (4 byte)

	guid, _ := xid.FromString(task.GetGUID())

	jsonByte, err := json.Marshal(&task)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	member := []byte(task.GetGUID())
	score1 := ledis.ScorePair{
		Score:  guid.Time().Unix(),
		Member: member, // can not store whole task, cause ledisdb has limit size (65K) in z/s/h
	}
	number, err := l.db.ZAdd(byteKeys, score1)
	if err != nil {
		log.Fatalf("DB Store error! %v", err)
	}
	err = l.db.Set(member, jsonByte) // has 1GB size limit
	if err != nil {
		l.db.ZRem(byteKeys, member) // rollback
		log.Fatalf("DB Store error! %v", err)
	}
	return string(number), nil
}

// Get members of a key
func (l LedisDB) Get(uid uint32, start int, stop int) *TaskCollection {
	byteKeys := (*[4]byte)(unsafe.Pointer(&uid))[:] // 32 bit id (4 byte)
	scorePairs, _ := l.db.ZRevRange(byteKeys, start, stop)
	tasks := make(map[int]*pb.Task)
	for index, scorePair := range scorePairs {
		score := scorePair.Score
		member := scorePair.Member
		value, _ := l.db.Get(member) // Get
		unixTimeUTC := time.Unix(score, 0)
		log.Debugf("Score: %v", unixTimeUTC.Format(time.RFC3339))
		task := &pb.Task{}
		json.Unmarshal(value, &task)
		log.Debugf("Member: %+v", string(task.GetOutput()))
		tasks[index] = task
	}
	return &TaskCollection{Tasks: tasks}
}

// SetTask add new task to the list of the tasks
func (l LedisDB) SetTask(task *pb.Task) (bool, error) {
	key := []byte("TaskList")
	jsonByte, jerr := json.Marshal(&task)
	if jerr != nil {
		log.Fatal("Json encode error:", jerr)
	}
	l.db.LClear(key)
	added, err := l.db.HSet(key, ledis.PutInt64(int64(task.UID)), jsonByte)
	if err != nil {
		log.Fatalf("DB HSet error! %v", err)
	}
	return (added == 1), nil
}

// GetTasks returns the list of the tasks
func (l LedisDB) GetTasks(from int32, limit int32) *TaskCollection {
	key := []byte("TaskList")
	list, err := l.db.HScan(key, ledis.PutInt64(int64(from)), int(limit), true, "")
	if err != nil {
		log.Fatalf("DB HScan error! %v", err)
	}
	tasksUnsorted := make(map[string]ledis.FVPair)
	var keys []string
	for _, fieldValue := range list {
		task := &pb.Task{}
		json.Unmarshal(fieldValue.Value, &task)
		keys = append(keys, task.GUID)
		tasksUnsorted[task.GUID] = fieldValue
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	tasks := make(map[int]*pb.Task)
	for index, guid := range keys {
		task := &pb.Task{}
		fvPair := tasksUnsorted[guid]
		json.Unmarshal(fvPair.Value, &task)
		tasks[index] = task
	}
	return &TaskCollection{Tasks: tasks}
}

// Lock create a lock
func (l LedisDB) Lock(key string, timeout int32) (bool, error) {
	db, erro := l.ledis.Select(1)
	byteKey := []byte(key)
	exists, err := db.Exists(byteKey)
	if err != nil || exists == 0 {
		err := db.Set(byteKey, []byte("1")) // value doesn't matter
		if err != nil {
			return false, err
		}
		a, errorr := db.Expire(byteKey, int64(timeout)) // expire time
		log.Warnf("Expires in :: %+v, %v , %v", a, errorr, timeout)
		return true, nil
	}
	log.Warnf("Lock: %+v, status: %+v : %+v", key, exists, erro)
	return (exists == 0), nil
}

// Release release lock
func (l LedisDB) Release(key string) (bool, error) {
	db, _ := l.ledis.Select(1)
	byteKey := []byte(key)
	if _, err := db.Del(byteKey); err != nil {
		return false, err
	}
	return true, nil
}

// Close members of a key
func (l LedisDB) Close() {
	l.ledis.Close()
}
