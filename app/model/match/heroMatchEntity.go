package match

import (
	heroModel "dotaapi/app/model/heroes"
	"dotaapi/app/service/heroes"
	"strconv"
)

//HeroMatchEntity  entity
type HeroMatchEntity struct {
	Games      float64           `json:"games"` //总场次
	HeroID     float64           `json:"heroId"`
	Hero       *heroModel.Entity `json:"hero"`
	Win        float64           `json:"win"`        //胜场
	LastPlayed float64           `json:"lastPlayed"` //最后一次使用时间
}

//CreateHeroMatchEntity create a match entity
func CreateHeroMatchEntity(item map[string]interface{}) *HeroMatchEntity {
	entity := new(HeroMatchEntity)
	entity.Games = item["games"].(float64)
	heroID, _ := strconv.ParseFloat(item["hero_id"].(string), 64)
	entity.HeroID = heroID
	entity.Win = item["win"].(float64)
	entity.LastPlayed = item["last_played"].(float64)

	heroes := heroes.GetHeroes()
	for _, v := range heroes {
		if v.ID == entity.HeroID {
			entity.Hero = v
		}
	}

	return entity
}

//BatchCreateHeroMatchEntity create a new hero match entity
func BatchCreateHeroMatchEntity(items []interface{}) []*HeroMatchEntity {
	matches := make([]*HeroMatchEntity, len(items))
	for k, v := range items {
		matches[k] = CreateHeroMatchEntity(v.(map[string]interface{}))
	}
	return matches
}
