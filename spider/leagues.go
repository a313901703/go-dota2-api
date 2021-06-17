package spider

import (
	"encoding/json"
	"fmt"

	// "encoding/json"
	// "regexp"
	// "strings"
	// "sync"
	leagueModel "dotaapi/app/model/league"
	"dotaapi/library/httprequest"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

//GetLeagues 赛事列表
func GetLeagues() {
	// data, _ := g.Redis().DoVar("GET", keyLeagues)
	// if data.String() != "" {
	// 	return
	// }
	params := make(map[string]interface{})
	params["page_size"] = "200"
	config := httprequest.Config{URL: leadguesURL, Data: params}
	resp := httprequest.FuncDataRequest(config)
	if resp == nil {
		g.Log().Error("sync leagues error")
		return
	}
	leagues := resp.([]interface{})
	leaguesArr := garray.NewArray(true)

	for _, v := range leagues {
		var leagueModel *leagueModel.Entity
		if err := gconv.Struct(v, &leagueModel); err != nil {
			panic(err)
		}
		src := fmt.Sprintf(leagueImgURL, gconv.String(leagueModel.LeagueID))
		if !fetchImg(src) {
			continue
		}
		// fmt.Println(src)
		leagueModel.Thumb = src
		leaguesArr.Append(leagueModel)
	}
	b, _ := json.Marshal(leaguesArr)
	g.Redis().DoVar("SET", keyLeagues, string(b))
	fmt.Println("sync league success")
}
