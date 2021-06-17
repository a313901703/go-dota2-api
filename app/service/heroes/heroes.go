package heroes

import (
	"dotaapi/app/model/heroes"
	"dotaapi/library/httprequest"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

const (
	keyHeroesList  = "redis_heroes_list"
	heroStatsURL   = "heroStats"
	keyHeroesStats = "redis_heroes_stats"
)

var heroURL string = "/fundata-dota2-free/v2/raw/hero"

//GetHeroes 从fundata api 获取英雄列表数据
//优先验证redis是否存在
func GetHeroes() []*heroes.Entity {
	v, err := g.Redis().DoVar("GET", keyHeroesList)
	var list []interface{}
	data := v.String()
	if err != nil || data == "" {
		config := httprequest.Config{URL: heroURL}
		resp := httprequest.FuncDataRequest(config)
		if resp != nil {
			list = resp.([]interface{})
			//转成字符串并存入redis
			jsonData, _ := json.Marshal(list)
			g.Redis().DoVar("SET", keyHeroesList, string(jsonData))
		}
	} else {
		//处理redis中的数据
		json.Unmarshal([]byte(data), &list)
	}
	var items = make([]*heroes.Entity, len(list))
	for k, v := range list {
		item := v.(map[string]interface{})
		items[k] = heroes.NewHeroEntity(item)
	}
	return items
}

//GetHeroesMap  map
func GetHeroesMap() map[float64]*heroes.Entity {
	s := GetHeroes()
	m := make(map[float64]*heroes.Entity)
	for _, v := range s {
		m[v.ID] = v
	}
	return m
}

//GetHero  hero 详情
func GetHero(ID string) (*gmap.AnyAnyMap, bool) {
	isOk := false
	s := GetHeroes()
	heroStats := getHeroStats()
	m := gmap.Map{}
	data, _ := g.Redis().DoVar("GET", "redis_heroes_info")
	json.Unmarshal([]byte(data.String()), &m)
	float, _ := strconv.ParseFloat(ID, 64)
	ret := gmap.New()
	for _, v := range s {
		if v.ID == float {
			info := m.Get(ID)
			heroStat := heroStats.Get(v.ID)
			ret.Set("info", info)
			ret.Set("stat", heroStat)
			ret.Set("base", v)
			//ret.Set("", value)
			isOk = true
		}
	}
	return ret, isOk
}

func getHeroStats() *gmap.AnyAnyMap {
	v, _ := g.Redis().DoVar("GET", keyHeroesStats)
	data := v.String()
	heroStatsMap := gmap.New()
	if data == "" {
		config := httprequest.Config{URL: heroStatsURL}
		resp := httprequest.OpenDotaRequest(config)
		//转成字符串并存入redis
		jsonData, _ := json.Marshal(resp)
		if string(jsonData) != "" {
			g.Redis().DoVar("SET", keyHeroesStats, string(jsonData))
		}
		data = string(jsonData)
	}

	if data != "" {
		var items []interface{}
		err := json.Unmarshal([]byte(data), &items)
		if err != nil {
			g.Log().Info(fmt.Sprintf("OpenDotaRequest JsonToMap err: %s \n", err))
		}
		for _, v := range items {
			var statEntity *heroes.StatEntity
			if err := gconv.Struct(v, &statEntity); err != nil {
				//fmt.Println(err)
			} else {
				heroStatsMap.Set(statEntity.ID, statEntity)
			}
		}
	}
	return heroStatsMap
}

//GetAbilities 技能信息
func GetAbilities() *gmap.AnyAnyMap {
	v, _ := g.Redis().DoVar("GET", "redis_ability")
	data := v.String()
	m := gmap.New()
	json.Unmarshal([]byte(data), &m)
	return m
}
