package match

import (
	heroModel "dotaapi/app/model/heroes"
	"dotaapi/app/service/heroes"
)

//Entity match entity
type Entity struct {
	Kills   float64 `json:"kills"`
	Deaths  float64 `json:"deaths"`
	Assists float64 `json:"assists"` //助攻
	// LastHits    float64 `json:"lastHits"`    //正补
	Duration float64 `json:"duration"` //时长
	GameMode float64 `json:"gameMode"` //比赛类型
	// GoldPerMin  float64 `json:"goldPerMin"`  //每分钟金钱
	// XpPerMin    float64 `json:"xpPerMin"`    //每分钟经验
	// HeroDamage  float64 `json:"heroDamage"`  //伤害
	// HeroHealing float64 `json:"HeroHealing"` //治疗
	// TowerDamage float64 `json:"towerDamage"` //防御塔伤害

	LobbyType  float64 `json:"lobbyType"`  //大厅类型
	MatchID    float64 `json:"matchID"`    //比赛ID
	PlayerSlot float64 `json:"playerSlot"` // 0-4天辉   其余夜魇
	StartTime  float64 `json:"startTime"`  //开始时间
	RadiantWin bool    `json:"radiantWin"` //天辉胜利
}

//PlayerEntity player match entity
type PlayerEntity struct {
	HeroID float64           `json:"heroId"`
	Hero   *heroModel.Entity `json:"hero"`
	Win    bool              `json:"win"` //胜利
	*Entity
}

//CreateNewEntity create a match entity
func CreateNewEntity(item map[string]interface{}) *Entity {
	entity := new(Entity)
	entity.Kills = item["kills"].(float64)
	entity.Deaths = item["deaths"].(float64)
	entity.Assists = item["assists"].(float64)
	//entity.LastHits = item["last_hits"].(float64)
	entity.Duration = item["duration"].(float64)
	entity.GameMode = item["game_mode"].(float64)
	// entity.GoldPerMin = item["gold_per_min"].(float64)
	// entity.XpPerMin = item["xp_per_min"].(float64)
	// entity.HeroDamage = item["hero_damage"].(float64)
	// entity.HeroHealing = item["hero_healing"].(float64)
	// entity.TowerDamage = item["tower_damage"].(float64)
	entity.LobbyType = item["lobby_type"].(float64)
	entity.MatchID = item["match_id"].(float64)
	entity.PlayerSlot = item["player_slot"].(float64)
	entity.StartTime = item["start_time"].(float64)
	entity.RadiantWin = item["radiant_win"].(bool)
	return entity
}

//CreateNewPlayerEntity create a new player match entity
func CreateNewPlayerEntity(item map[string]interface{}) *PlayerEntity {
	entity := CreateNewEntity(item)
	playerEntity := &PlayerEntity{Entity: entity}
	playerEntity.HeroID = item["hero_id"].(float64)
	isRadiant := true
	if playerEntity.PlayerSlot > 4 {
		isRadiant = false
	}
	playerEntity.Win = (entity.RadiantWin == isRadiant)
	heroes := heroes.GetHeroes()
	for _, v := range heroes {
		if v.ID == playerEntity.HeroID {
			playerEntity.Hero = v
		}
	}
	return playerEntity
}

//BatchCreateNewPlayerEntity create a new player match entity
func BatchCreateNewPlayerEntity(items []interface{}) []*PlayerEntity {
	matches := make([]*PlayerEntity, len(items))
	for k, v := range items {
		matches[k] = CreateNewPlayerEntity(v.(map[string]interface{}))
	}
	return matches
}
