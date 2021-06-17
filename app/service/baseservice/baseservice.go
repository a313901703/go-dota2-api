package baseservice

import (
	"dotaapi/library/httprequest"
	"encoding/json"
	"errors"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
)

//Service  service基类
type Service struct {
	URL  string
	Key  string //redis key
	From int64  // 1. opendota  2. fundata  3. steam
	Exp  int64  //过期时间

	Page     int
	PageSize int
}

const (
	opendota = 1
	fundata  = 2
)

//Get 统一处理players数据
//先从redis hash中获取数据  如果数据不存在  调取API
func (s Service) Get() (resp interface{}, err error) {
	data := ""
	if s.Key != "" {
		data = GetRedis(s.Key)
	}

	if data == "" {
		config := httprequest.Config{URL: s.URL}
		switch s.From {
		case fundata:
			resp = httprequest.FuncDataRequest(config)
		default:
			resp = httprequest.OpenDotaRequest(config)
		}

		if resp == nil {
			err = errors.New("请求失败")
			return
		}

		//转成字符串并存入redis
		jsonData, _ := json.Marshal(resp)
		if s.Key != "" {
			if s.Exp > 0 {
				SetRedisWithExpire(s.Key, string(jsonData), s.Exp)
			} else {
				SetRedis(s.Key, string(jsonData))
			}
		}

	} else {
		//处理redis中的数据
		json.Unmarshal([]byte(data), &resp)
	}
	// fmt.Println("resp", resp)
	return
}

//Pagination 统一处理分页
func (s Service) Pagination(items *garray.Array) *gmap.AnyAnyMap {
	var ret []interface{}

	pagination := gmap.New()
	pagination.Set("page", s.Page)
	pagination.Set("pageSize", s.PageSize)

	m := gmap.New()
	m.Set("pagination", pagination)
	m.Set("data", &ret)

	if items.Len() == 0 {
		pagination.Set("count", 0)
		return m
	}
	pagination.Set("count", items.Len())
	if s.Page > 0 && s.PageSize > 0 {
		end := (s.Page * s.PageSize)
		start := (s.Page - 1) * s.PageSize
		ret = items.Range(start, end)
	} else {
		ret = items.Range(0)
	}
	return m
}
