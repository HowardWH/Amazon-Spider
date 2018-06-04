/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package myredis

import (
	"errors"
	"time"

	"github.com/hunterhug/parrot/util"
	"gopkg.in/redis.v4"
)

// redis tool

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type MyRedis struct {
	Config RedisConfig
	Client *redis.Client
}

// return myredis
func NewRedis(config RedisConfig) (*MyRedis, error) {
	myredis := &MyRedis{Config: config}
	client := redis.NewClient(&redis.Options{
		Addr:        config.Host,
		Password:    config.Password, // no password set
		DB:          config.DB,       // use default DB
		MaxRetries:  5,               // fail command retry 2
		PoolSize:    40,              // redis pool size
		DialTimeout: util.Second(20),
		// another options is default
	})

	pong, err := client.Ping().Result()
	if err == nil && pong == "PONG" {
		myredis.Client = client
	}
	return myredis, err
}

func NewRedisPool(config RedisConfig, size int) (*MyRedis, error) {
	myredis := &MyRedis{Config: config}
	client := redis.NewClient(&redis.Options{
		Addr:        config.Host,
		Password:    config.Password, // no password set
		DB:          config.DB,       // use default DB
		MaxRetries:  5,               // fail command retry 2
		PoolSize:    size,            // redis pool size
		DialTimeout: util.Second(20),
		// another options is default
	})

	pong, err := client.Ping().Result()
	if err == nil && pong == "PONG" {
		myredis.Client = client
	}
	return myredis, err
}

// set key
func (db *MyRedis) Set(key string, value string, expire time.Duration) error {
	return db.Client.Set(key, value, expire).Err()
}

// get key
func (db *MyRedis) Get(key string) (string, error) {
	result, err := db.Client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("redis key does not exists")
	} else if err != nil {
		return "", err
	} else {
		return result, err
	}
}

func (db *MyRedis) Lpush(key string, values ...interface{}) (int64, error) {
	return db.Client.LPush(key, values...).Result()
}

func (db *MyRedis) Lpushx(key string, values interface{}) (int64, error) {
	num, err := db.Client.LPushX(key, values).Result()
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, errors.New("Redis List not exist")
	} else {
		return num, err
	}
}

func (db *MyRedis) Rpush(key string, values ...interface{}) (int64, error) {
	return db.Client.RPush(key, values...).Result()
}

func (db *MyRedis) Rpushx(key string, values interface{}) (int64, error) {
	num, err := db.Client.RPushX(key, values).Result()
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, errors.New("Redis List not exist")
	} else {
		return num, err
	}
}

func (db *MyRedis) Llen(key string) (int64, error) {
	return db.Client.LLen(key).Result()
}

func (db *MyRedis) Hlen(key string) (int64, error) {
	return db.Client.HLen(key).Result()
}

func (db *MyRedis) Rpop(key string) (string, error) {
	return db.Client.RPop(key).Result()
}

func (db *MyRedis) Lpop(key string) (string, error) {
	return db.Client.LPop(key).Result()
}

func (db *MyRedis) Brpop(timeout int, keys ...string) ([]string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BRPop(timeouts, keys...).Result()
}

// if timeout is zero will be block until...
// and if  keys has many will return one such as []string{"pool","b"},pool is list,b is value
func (db *MyRedis) Blpop(timeout int, keys ...string) ([]string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BLPop(timeouts, keys...).Result()
}

func (db *MyRedis) Brpoplpush(source, destination string, timeout int) (string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BRPopLPush(source, destination, timeouts).Result()
}

func (db *MyRedis) Rpoplpush(source, destination string) (string, error) {
	return db.Client.RPopLPush(source, destination).Result()
}

func (db *MyRedis) Hexists(key, field string) (bool, error) {
	return db.Client.HExists(key, field).Result()
}

func (db *MyRedis) Hget(key, field string) (string, error) {
	return db.Client.HGet(key, field).Result()
}

func (db *MyRedis) Hset(key, field, value string) (bool, error) {
	return db.Client.HSet(key, field, value).Result()
}

// return item rem number if count==0 all rem if count>0 from the list head to rem
func (db *MyRedis) Lrem(key string, count int64, value interface{}) (int64, error) {
	return db.Client.LRem(key, count, value).Result()
}
