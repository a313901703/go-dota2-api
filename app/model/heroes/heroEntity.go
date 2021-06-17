package heroes

import (
	"strings"

	"github.com/gogf/gf/frame/g"
)

//Entity  英雄实例
type Entity struct {
	ID          float64       `json:"id"`
	Legs        float64       `json:"legs"`
	CnName      string        `json:"cnName"`
	EnName      string        `json:"enName"`
	Name        string        `json:"name"`
	AttackType  string        `json:"attackType"`
	PrimaryAttr string        `json:"primaryAttr"`
	Roles       []interface{} `json:"roles"`
	Thumb       string        `json:"thumb"`
}

//InfoEntity  详情
type InfoEntity struct {
	*Entity
	Ability string `json:"ability"`
	Counts  string `json:"counts"`
	Src     string `json:"src"`
	Talents string `json:"talents"`
	WinRate string `json:"winRate"`
}

//StatEntity  属性
type StatEntity struct {
	ID              float64 `json:"id"`
	AgiGain         float64 `json:"agiGain"`     //敏捷成长
	IntGain         float64 `json:"intGain"`     //智力成长
	StrGain         float64 `json:"strGain"`     //力量成长
	AttackRange     float64 `json:"attackRange"` //攻击距离
	AttackRate      float64 `json:"attackRate"`  //攻击速率
	AttackType      string  `json:"attackType"`
	BaseAgi         float64 `json:"baseAgi"`         //基础敏捷
	BaseStr         float64 `json:"baseStr"`         //基础力量
	BaseInt         float64 `json:"baseInt"`         //基础智力
	BaseArmor       float64 `json:"baseArmor"`       //基础护甲
	BaseHealth      float64 `json:"baseHealth"`      //基础血量
	BaseHealthRegen float64 `json:"baseHealthRegen"` //基础恢复
	BaseMana        float64 `json:"baseMana"`        //基础魔法
	BaseManaRegen   float64 `json:"baseManaRegen"`   //基础回魔
	BaseAttackMax   float64 `json:"baseAttackMax"`   //基础最大攻击
	BaseAttackMin   float64 `json:"baseAttackMin"`   //基础最小攻击
	MoveSpeed       float64 `json:"moveSpeed"`       //移动速度
	TurnRate        float64 `json:"turnRate"`        //转身速度
}

//NewHeroEntity  hero实例
func NewHeroEntity(item map[string]interface{}) *Entity {
	entity := new(Entity)
	entity.ID = item["hero_id"].(float64)
	entity.Legs = item["legs"].(float64)
	entity.CnName = item["cn_name"].(string)
	entity.EnName = item["en_name"].(string)
	entity.Name = strings.TrimPrefix(item["hero_name"].(string), "npc_dota_hero_")
	entity.AttackType = item["attack_type"].(string)
	entity.PrimaryAttr = item["primary_attr"].(string)
	entity.Roles = item["cn_roles"].([]interface{})
	name := strings.TrimPrefix(item["hero_name"].(string), "npc_dota_hero_")
	entity.Thumb = g.Cfg().Get("params.imgDomain").(string) + name
	return entity
}
