package team

import (
	"dotaapi/app/service/baseservice"
	heroesService "dotaapi/app/service/heroes"
	"errors"
	"fmt"
	"sync"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
)

//Service 继承baseservice
type Service struct {
	baseservice.Service
}

const (
	teamsURL       = "/teams"
	teamURL        = "/teams/%s"
	teamMatchesURL = "/teams/%s/matches"
	teamPlayersURL = "/teams/%s/players"
	teamHeroesURL  = "/teams/%s/heroes"
)

const (
	redisKeyAll     = "teams"
	redisKeyOne     = "teams_%s"
	redisKeyPlayers = "teams_players_%s"
	redisKeyHeroes  = "teams_heroes_%s"
)

//GetAll 获取全部team
func (s Service) GetAll() *gmap.AnyAnyMap {
	s.URL = teamsURL
	s.Key = redisKeyAll
	s.Exp = 3600 * 24
	resp, err := s.Get()

	teams := garray.NewArray()
	if err != nil {
		return gmap.New()
	}
	items := resp.([]interface{})
	teams.SetArray(items)
	return s.Pagination(teams)
}

//GetInfo team 信息  成员  常用hero信息
func (s Service) GetInfo(id string) (*gmap.AnyAnyMap, error) {
	m := gmap.New()
	var err error
	var wg sync.WaitGroup
	wg.Add(3)

	var item, players, heroes interface{}
	go func() {
		item, err = s.GetOne(id)
		wg.Done()
	}()

	go func() {
		players, err = s.GetPlayers(id)
		wg.Done()
	}()

	go func() {
		heroes, err = s.GetHeroes(id)
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		return m, err
	}
	for k, v := range item.(map[string]interface{}) {
		m.Set(k, v)
	}
	m.Set("heroes", heroes)
	m.Set("players", players)

	return m, err
}

//GetOne 获取单个team
func (s Service) GetOne(id string) (item interface{}, err error) {
	if id == "" {
		err = errors.New("无效的team_id")
		return
	}
	s.URL = fmt.Sprintf(teamURL, id)
	s.Key = fmt.Sprintf(redisKeyOne, id)
	s.Exp = 3600 * 24
	item, err = s.Get()
	if err != nil || item == nil {
		err = errors.New("未获取到队伍信息")
	}
	return
}

//GetMatches 获取单个team的比赛
func (s Service) GetMatches(id string) (items interface{}, err error) {
	if id == "" {
		err = errors.New("无效的team_id")
		return
	}
	s.URL = fmt.Sprintf(teamMatchesURL, id)
	items, err = s.Get()
	if err != nil {
		return
	}
	if items == nil {
		err = errors.New("未获取到队伍信息")
	}
	return
}

//GetPlayers 获取单个team的选手
func (s Service) GetPlayers(id string) (items interface{}, err error) {
	if id == "" {
		err = errors.New("无效的team_id")
		return
	}
	s.URL = fmt.Sprintf(teamPlayersURL, id)
	s.Key = fmt.Sprintf(redisKeyPlayers, id)
	s.Exp = 3600 * 24

	items, err = s.Get()
	if err != nil {
		return
	}
	if items == nil {
		err = errors.New("未获取到队伍信息")
	}
	return
}

//GetHeroes 获取team heroes
func (s Service) GetHeroes(id string) ([]interface{}, error) {
	var err error
	var array []interface{}
	if id == "" {
		err = errors.New("无效的team_id")
		return array, err
	}
	s.URL = fmt.Sprintf(teamHeroesURL, id)
	s.Key = fmt.Sprintf(redisKeyHeroes, id)
	s.Exp = 3600 * 24

	items, err := s.Get()
	if err != nil || items == nil {
		err = errors.New("未获取到队伍信息")
		return array, err
	}
	array = items.([]interface{})
	array = array[0:10]
	heroMaps := heroesService.GetHeroesMap()
	for k, v := range array {
		hero := v.(map[string]interface{})
		heroID := hero["hero_id"].(float64)
		hero["hero"] = heroMaps[heroID]
		array[k] = hero
	}

	return array, err
}
