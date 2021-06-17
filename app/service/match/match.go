package match

import (
	"dotaapi/app/model/match"
	"dotaapi/app/service/baseservice"
	"dotaapi/app/service/heroes"
	"dotaapi/app/service/items"
	"dotaapi/library/httprequest"
	"fmt"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/encoding/gjson"
)

//Service 实例
type Service struct {
	baseservice.Service
}

const (
	matchInfoURL = "matches/%s"
)

//FindOne 单场比赛的详细数据
func (s Service) FindOne(matchID string) *match.InfoEntity {
	uri := fmt.Sprintf(matchInfoURL, matchID)
	config := httprequest.Config{URL: uri}
	resp := httprequest.OpenDotaRequestAsStr(config)
	item := new(match.InfoEntity)
	if resp != "" {
		if j, err := gjson.DecodeToJson(resp); err != nil {
			panic(err)
		} else {
			if err := j.ToStruct(item); err != nil {
				panic(err)
			}
		}
		//英雄数据
		heroesMap := heroes.GetHeroesMap()
		//item数据
		itemsMap := items.GetItemsMap()
		//技能信息
		abilities := heroes.GetAbilities()
		//transform
		for _, v := range item.Players {
			v.Hero = heroesMap[v.HeroID]
			items := []float64{v.Item0, v.Item1, v.Item2, v.Item3, v.Item4, v.Item5}
			backpacks := []float64{v.Backpack0, v.Backpack1, v.Backpack2}
			for _, itemID := range items {
				v.Items = append(v.Items, itemsMap[itemID])
			}
			for _, backpackID := range backpacks {
				v.Backpacks = append(v.Backpacks, itemsMap[backpackID])
			}
			heroAbilities := make([]interface{}, len(v.AbilityUpgradesArr))
			for k, abilityID := range v.AbilityUpgradesArr {
				heroAbilities[k] = abilities.Get(gconv.String(abilityID))
			}
			v.Abilities = heroAbilities
		}

	}
	return item
}
