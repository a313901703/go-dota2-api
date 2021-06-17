package news

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
}

//GetNews 获取新闻列表
func (s Service) GetNews() *gmap.AnyAnyMap {
	v, _ := g.Redis().DoVar("GET", "redis_news_list")
	data := v.String()

	var news *garray.Array

	if data != "" {
		json.Unmarshal([]byte(data), &news)
	}
	return s.Pagination(news)
}
