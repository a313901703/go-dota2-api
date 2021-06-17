package players

import (
	"dotaapi/app/model/match"
	"dotaapi/app/model/players"
	"dotaapi/library/httprequest"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/gcache"
)

type requestConfig struct {
	uri      string
	key      string
	playerID string // 玩家ID
	apiType  int64  // 1. opendota  2. fundata  3. steam
}

const (
	openDotaAPIType = 1
	fundataAPIType  = 2
	steamAPIType    = 3
)

//opendota api
const (
	playerInfoURL = "/players/%s"
	wlURL         = "/players/%s/wl"
	recentGameURL = "/players/%s/recentMatches"
	peersURL      = "/players/%s/peers"
	countURL      = "/players/%s/totals"
	userdHeroURL  = "/players/%s/heroes"
	matchURL      = "/players/%s/matches"
)

//GetInfo  详情、队友、近期比赛、win & lost 等数据
// @param   id         string     玩家ID
// @param   expects    []string   期望获取的资源   info、wl、recentMatches、peers、count、usedHero
func GetInfo(id string, expects []string) *gmap.Map {
	if len(expects) == 0 {
		expects = []string{"info", "wl", "recentMatches", "peers", "usedHero", "count"}
	}
	mapData := gmap.New()
	var wg sync.WaitGroup
	wg.Add(len(expects))

	for _, v := range expects {
		switch v {
		case "info":
			go func(v string) {
				data, _ := Info(id)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		case "wl":
			go func(v string) {
				data, _ := Wl(id)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		case "recentMatches":
			go func(v string) {
				data, _ := Matches(id, 10, 1)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		case "peers":
			go func(v string) {
				data, _ := Peers(id)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		case "count":
			go func(v string) {
				data := Count(id)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		case "usedHero":
			go func(v string) {
				data := UsedHero(id)
				mapData.Set(v, data)
				wg.Done()
			}(v)
		}

	}
	wg.Wait()
	return mapData
}

//Info 获取玩家详情
func Info(id string) (item *players.InfoEntity, err error) {
	uri := fmt.Sprintf(playerInfoURL, id)
	requestConfig := requestConfig{uri: uri, key: "info", playerID: id}
	resp, err := requestConfig.get()
	fmt.Println("resp", resp.(map[string]interface{}))
	if resp != nil {
		data := resp.(map[string]interface{})
		item = players.NewPlayerInfoEntity(data)
	}
	return item, err
}

//Wl win && lose
func Wl(id string) (item map[string]interface{}, err error) {
	uri := fmt.Sprintf(wlURL, id)
	requestConfig := requestConfig{uri: uri, key: "wl", playerID: id}
	resp, err := requestConfig.get()
	if err == nil {
		item = resp.(map[string]interface{})
	}
	return
}

//RecentMatches 近期比赛
func RecentMatches(id string) (matches []*match.PlayerEntity, err error) {
	uri := fmt.Sprintf(recentGameURL, id)
	config := httprequest.Config{URL: uri}
	resp := httprequest.OpenDotaRequest(config)
	if resp != nil {
		items := resp.([]interface{})
		matches = match.BatchCreateNewPlayerEntity(items)
	}
	return matches, err
}

//Peers 队友
func Peers(id string) (friends []*players.FriendInfoEntity, err error) {
	uri := fmt.Sprintf(peersURL, id)
	requestConfig := requestConfig{uri: uri, key: "peers", playerID: id}
	var items interface{}
	items, err = requestConfig.get()
	if items != nil {
		friends = players.BatchCreateFriendInfoEntity(items.([]interface{}))
	}
	return friends, err
}

//get 统一处理players数据
//先从redis hash中获取数据  如果数据不存在  调取API
func (r requestConfig) get() (resp interface{}, err error) {
	redis := PlayerRedis{id: r.playerID}
	data := redis.Get(r.key)

	if data == "" {
		// URL := fmt.Sprintf(playerInfoURL,id)
		config := httprequest.Config{URL: r.uri}
		switch r.apiType {
		case fundataAPIType:
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
		redis.Set(r.key, string(jsonData))
	} else {
		//处理redis中的数据
		json.Unmarshal([]byte(data), &resp)
	}
	// fmt.Println("resp", resp)
	return
}

//UsedHero 常用hero
func UsedHero(id string) (heroes []*match.HeroMatchEntity) {
	key := "used_heroes_" + id

	resp := gcache.Get(key)
	if resp == nil {
		uri := fmt.Sprintf(userdHeroURL, id)
		config := httprequest.Config{URL: uri}
		resp = httprequest.OpenDotaRequest(config)
		gcache.Set(key, resp, 3600*time.Second)
	}
	if resp != nil {
		items := resp.([]interface{})
		heroes = match.BatchCreateHeroMatchEntity(items)
	}
	return heroes
}

//Count 统计
func Count(id string) map[string]interface{} {
	key := "player_total_" + id
	resp := gcache.Get(key)

	if resp == nil {
		uri := fmt.Sprintf(countURL, id)
		config := httprequest.Config{URL: uri}
		resp = httprequest.OpenDotaRequest(config)
		gcache.Set(key, resp, 3600*time.Second)
	}
	items := map[string]interface{}{}
	if resp != nil {
		counts := resp.([]interface{})
		for _, v := range counts {
			count := v.(map[string]interface{})
			items[count["field"].(string)] = count["sum"]
		}
	}

	return items
}

//Matches 比赛数据
func Matches(id string, limit int, page int) (matches []*match.PlayerEntity, err error) {
	uri := fmt.Sprintf(matchURL, id)
	paramsStr := ""
	if limit > 0 {
		paramsStr += "limit=" + strconv.Itoa(limit) + "&"
	}
	if page > 0 {
		offset := (page - 1) * limit
		paramsStr += "offset=" + strconv.Itoa(offset)
	}
	if paramsStr != "" {
		paramsStr = "?" + paramsStr
	}
	uri += paramsStr
	config := httprequest.Config{URL: uri}
	resp := httprequest.OpenDotaRequest(config)
	if resp != nil {
		items := resp.([]interface{})
		matches = match.BatchCreateNewPlayerEntity(items)
	}
	return matches, err
}
