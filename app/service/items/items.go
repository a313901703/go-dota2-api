package items

import (
	itemModel "dotaapi/app/model/item"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
)

const (
	keyList = "redis_items_list"
)

var itemURL string = "/fundata-dota2-free/v2/raw/item"

//Fetch 从fundata api 获取物品列表数据
//优先验证redis是否存在
func Fetch() *garray.Array {
	v, err := g.Redis().DoVar("GET", keyList)
	items := garray.NewArray()
	data := v.String()
	if err != nil || data == "" {
		// config := httprequest.Config{URL: itemURL}
		// list = httprequest.FuncDataRequest(config).([]interface{})
		// //转成字符串并存入redis
		// jsonData, _ := json.Marshal(list)
		// g.Redis().DoVar("SET", keyList, string(jsonData))
	} else {
		json.Unmarshal([]byte(data), &items)
	}
	return items
}

//GetItemsMap  map
func GetItemsMap() map[float64]*itemModel.InfoEntity {
	s := Fetch()
	m := make(map[float64]*itemModel.InfoEntity)
	s.Iterator(func(k int, v interface{}) bool {
		var item *itemModel.InfoEntity
		if err := gconv.StructDeep(v, &item); err != nil {
			fmt.Println("error:", v)
		} else {
			m[item.ID] = item
		}
		return true
	})
	return m
}
