package cache

import (
	"fmt"
	red "github.com/gomodule/redigo/redis"
	"time"
	"github.com/astaxie/beego"
)

type Redis struct {
	pool     *red.Pool
	keySpace string
}

func NewRedis(pri string) *Redis {
	redis := new(Redis)
	redis.keySpace = beego.AppConfig.String(pri + "keySpace")
	maxIdle,_ := beego.AppConfig.Int(pri + "maxIdle")
	idleTimeout,_ := beego.AppConfig.Int(pri + "idleTime")
	idx,_ := beego.AppConfig.Int(pri + "idx")
	readTimeout,_ := beego.AppConfig.Int(pri + "readTimeout")
	writeTimeout,_ := beego.AppConfig.Int(pri + "writeTimeout")
	connectTimeout,_ := beego.AppConfig.Int(pri + "connectTimeout")
	redis.pool = &red.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   0,
		IdleTimeout: time.Duration(idleTimeout),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				beego.AppConfig.String(pri + "host"),
				red.DialReadTimeout(time.Duration(readTimeout)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(writeTimeout)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(connectTimeout)*time.Millisecond),
				red.DialDatabase(idx),
				red.DialPassword(beego.AppConfig.String(pri + "password")),
			)
		},
	}
	return redis
}

func (this *Redis) GetKey(key string) string {
	if this.keySpace != "" {
		key = fmt.Sprintf("%s:%s", this.keySpace, key)
	}
	return key
}

func (this *Redis) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := this.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	parmas := make([]interface{}, 0)
	switch key := key.(type) {
	case string:
		parmas = append(parmas, this.GetKey(key))
	case []string:
		for _, v := range key {
			parmas = append(parmas, this.GetKey(v))
		}
	}

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

func (this *Redis) ExecParams(cmd string, args ...interface{}) (interface{}, error) {
	con := this.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	return con.Do(cmd, args...)
}
