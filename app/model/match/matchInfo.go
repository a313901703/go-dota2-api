package match

import (
	heroModel "dotaapi/app/model/heroes"
	itemModel "dotaapi/app/model/item"
)

//InfoEntity  比赛详情
type InfoEntity struct {
	Duration       float64             `json:"duration"`       //持续时长
	StartTime      float64             `json:"startTime"`      //开始时间
	FirstBloodTime float64             `json:"firstBloodTime"` //一血时间
	GameMode       float64             `json:"gameMode"`       //游戏类型
	MatchID        float64             `json:"matchID"`        //比赛ID
	MatchSeqNum    float64             `json:"matchSeqNum"`    //序列号  唯一
	RadiantWin     bool                `json:"radiantWin"`     //天辉胜利
	Players        []*InfoPlayerEntity `json:"players"`        //玩家数据
	//PicksBans      float64 `json:"picksBans"`      //选人
	// RadiantXpAdv   float64 `json:"radiantXpAdv"`   //天辉平均经验
	// RadiantGoldAdv float64 `json:"radiantGoldAdv"` //天辉平均金钱
}

//InfoPlayerEntity 比赛玩家信息
type InfoPlayerEntity struct {
	//玩家信息
	AccountID   float64 `json:"accountID"`   //玩家ID
	Personaname string  `json:"personaname"` //玩家名称
	IsRadiant   bool    `json:"isRadiant"`   //是否是天辉
	//英雄
	HeroID             float64           `json:"heroID"` //英雄ID
	Hero               *heroModel.Entity `json:"hero"`
	AbilityUpgradesArr []float64         `json:"abilityUpgradesArr"` //技能升级情况
	Abilities          []interface{}     `json:"abilities"`          //技能
	//数据
	Kills       float64 `json:"kills"`       //击杀
	Deaths      float64 `json:"deaths"`      //死亡
	Assists     float64 `json:"assists"`     //助攻
	LastHits    float64 `json:"lastHits"`    //正补
	Denies      float64 `json:"denies"`      //反补
	Kda         float64 `json:"kda"`         //kda
	GoldPerMin  float64 `json:"goldPerMin"`  //每分钟金钱
	TotalGold   float64 `json:"totalGold"`   //财产总和
	XpPerMin    float64 `json:"xpPerMin"`    //每分钟经验
	TotalXp     float64 `json:"totalXp"`     //经济总和
	HeroDamage  float64 `json:"heroDamage"`  //英雄伤害
	TowerDamage float64 `json:"towerDamage"` //推塔伤害
	HeroHealing float64 `json:"heroHealing"` //英雄治疗
	Level       float64 `json:"level"`       //英雄等级
	//物品
	Backpack0 float64 `json:"backpack0"` //背包
	Backpack1 float64 `json:"backpack1"` //背包
	Backpack2 float64 `json:"backpack2"` //背包
	//Backpack3 float64 `json:"backpack3"` //背包
	Backpacks []*itemModel.InfoEntity `json:"backpacks"` //背包
	Item0     float64                 `json:"item0"`     //装备
	Item1     float64                 `json:"item1"`     //装备
	Item2     float64                 `json:"item2"`     //装备
	Item3     float64                 `json:"item3"`     //装备
	Item4     float64                 `json:"item4"`     //装备
	Item5     float64                 `json:"item5"`     //装备
	Items     []*itemModel.InfoEntity `json:"items"`     //装备
}
