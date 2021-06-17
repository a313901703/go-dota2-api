package players

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
)

//PlayerRedis struct
type PlayerRedis struct {
	id string
}

const (
	keyPlayer = "redis_players_"
)

//Set  设置单个key
func (p PlayerRedis) Set(field string, value string) bool {
	redisKey := keyPlayer + p.id
	_, err := g.Redis().Do("HSET", redisKey, field, value)
	if err != nil {
		//set redis error
		return false
	}
	return true
}

//Get  获取单个key
func (p PlayerRedis) Get(field string) string {
	data, ok := p.GetAll()
	if ok == false || field == "" || data[field] == nil {
		return ""
	}
	value := data[field]
	return value.(string)
}

//GetAll  获取全部缓存
func (p PlayerRedis) GetAll() (data map[string]interface{}, ok bool) {
	var (
		err    error
		result *gvar.Var
	)

	redisKey := keyPlayer + p.id
	result, err = g.Redis().DoVar("HGETALL", redisKey)
	if err != nil {
		ok = false
		return
	}
	data = result.Map()
	ok = true
	return
}
