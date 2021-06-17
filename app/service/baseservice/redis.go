package baseservice

import (
	"github.com/gogf/gf/frame/g"
)

const redisPrefix = "redis_"

//SetRedis  设置redis数据
func SetRedis(k string, v string) bool {
	if k == "" || v == "" {
		// err = errors.New("key || value can not be blank")
		return false
	}
	k = redisPrefix + k
	g.Redis().DoVar("SET", k, v)
	return true
}

//GetRedis  获取redis数据
func GetRedis(k string) string {
	k = redisPrefix + k
	v, err := g.Redis().DoVar("GET", k)
	if err != nil {
		return ""
	}
	return v.String()
}

//SetRedisWithExpire 设置有过期时间的key
func SetRedisWithExpire(k string, v string, exp int64) bool {
	if SetRedis(k, v) {
		g.Redis().DoVar("EXPIRE", k, exp)
		return true
	}
	return false
}
