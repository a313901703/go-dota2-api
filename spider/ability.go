package spider

import (
	"dotaapi/library/httprequest"
	"encoding/json"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

//GetAbility 获取英雄技能
func GetAbility() {
	// data, _ := g.Redis().DoVar("GET", keyAbility)
	// if data.String() != "" {
	// 	return
	// }
	config := httprequest.Config{URL: abilityURL}
	abilityResp := httprequest.OpenDotaRequest(config)

	config.URL = abilitiesIdsURL
	idsResp := httprequest.OpenDotaRequest(config)
	ability := abilityResp.(map[string]interface{})
	ids := idsResp.(map[string]interface{})

	abilityMap := gmap.New()
	for k, v := range ids {
		if ability[v.(string)] != nil {
			item := ability[v.(string)].(map[string]interface{})
			if item["img"] != nil {
				item["img"] = "https://api.opendota.com" + item["img"].(string)
			}

			abilityMap.Set(k, item)
		}
	}
	b, _ := json.Marshal(abilityMap)
	g.Redis().DoVar("SET", keyAbility, string(b))
}
