package item

import (
	"strings"
)

//Entity  英雄实例
type Entity struct {
	Cost       float64 `json:"cost"`
	ID         float64 `json:"id"`
	CnName     string  `json:"cnName"`
	EnName     string  `json:"enName"`
	Name       string  `json:"name"`
	Recipe     float64 `json:"recipe"`     //是否是卷轴
	SecretShop float64 `json:"secretShop"` //是否可在神秘商店购买
	Thumb      string  `json:"thumb"`
	WinRate    string  `json:"winRate"`
	Counts     string  `json:"counts"`
	Desc       string  `json:"desc"`
	Elements   string  `json:"elements"`
}

//InfoEntity  英雄实例
type InfoEntity struct {
	Entity
	WinRate  string `json:"winRate"`
	Counts   string `json:"counts"`
	Desc     string `json:"desc"`
	Elements string `json:"elements"`
}

//NewEntity  hero实例
func NewEntity(item map[string]interface{}) Entity {
	var entity Entity
	//entity := new(Entity)
	entity.ID = item["id"].(float64)
	entity.Cost = item["cost"].(float64)
	entity.CnName = item["cn_name"].(string)
	entity.EnName = item["en_name"].(string)
	entity.Name = strings.TrimPrefix(item["name"].(string), "item_")
	entity.Recipe = item["recipe"].(float64)
	entity.SecretShop = item["secret_shop"].(float64)
	//entity.Thumb = g.Cfg().Get("params.imgDomain").(string) + "itgems/" + entity.Name + ".png"
	return entity
}
