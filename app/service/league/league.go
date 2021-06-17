package league

import (
	"dotaapi/app/service/baseservice"
	"encoding/json"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

//Service 实例
type Service struct {
	baseservice.Service
	// Page     int
	// PageSize int
}

//GetLeagues 赛事列表
func (s Service) GetLeagues() *gmap.AnyAnyMap {
	v, _ := g.Redis().DoVar("GET", "redis_leagues_list")
	data := v.String()

	var leagues *garray.Array

	if data != "" {
		json.Unmarshal([]byte(data), &leagues)
	}
	return s.Pagination(leagues)
}
