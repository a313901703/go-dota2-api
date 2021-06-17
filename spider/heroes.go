package spider

import (
	"dotaapi/app/model/heroes"
	heroService "dotaapi/app/service/heroes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

//GetHeroInfo 获取英雄头像，图片，技能等数据
func GetHeroInfo() {
	// data, _ := g.Redis().DoVar("GET", keyHeroesInfo)
	// if data.String() != "" {
	// 	return
	// }
	s := heroService.GetHeroes()
	m := make(map[string]*heroes.Entity)
	for _, v := range s {
		m[v.Name] = v
	}
	//图片数据
	spider := Spider{URL: maxBaseURL + "hero/rate/"}
	doc := spider.GetDoc()

	heroItemMaps := gmap.New(true)
	var wg sync.WaitGroup

	//获取头像  胜率  使用场次
	doc.Find(".table-list tbody tr").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		heroItemMap := make(map[string]string)
		name := ""
		s.Find("td").Each(func(i int, td *goquery.Selection) {

			//头像
			if i == 0 {
				src, _ := td.Find(".hero-img-list").First().Attr("src")
				heroItemMap["src"] = src
				name = strings.TrimPrefix(src, "http://cdn.dota2.com/apps/dota2/images/heroes/")
				name = strings.TrimSuffix(name, "_hphover.png")
			}
			//胜率
			if i == 1 {
				heroItemMap["winRate"] = td.Find("div").First().Text()
			}
			//场次
			if i == 2 {
				heroItemMap["counts"] = td.Find("div").First().Text()
			}
		})
		if m[name] == nil {
			go func(name string) {
				fmt.Println(name)
				wg.Done()
			}(name)
		} else {
			//英雄详情页面
			go func(name string, ID float64) {
				heroItemMap["talents"], heroItemMap["ability"] = parseHeroInfo(name, m[name].ID)
				heroItemMaps.Set(m[name].ID, heroItemMap)
				wg.Done()
			}(name, m[name].ID)
		}
	})
	wg.Wait()
	herosJSON, _ := json.Marshal(heroItemMaps)
	g.Redis().DoVar("SET", keyHeroesInfo, string(herosJSON))
	fmt.Println("sync heroes info down")
}

func parseHeroInfo(heroName string, ID float64) (string, string) {
	spider := Spider{URL: maxBaseURL + "hero/detail/" + heroName + "/"}
	doc := spider.GetDoc()
	talents := gmap.New()
	//天赋
	doc.Find(".talent_self").Each(func(i int, s *goquery.Selection) {
		left := strings.TrimSpace(s.Find(".talent_text_left").First().Text())
		right := strings.TrimSpace(s.Find(".talent_text_right").First().Text())
		talents.Set(10+i*5, left+","+right)
	})
	talentsJSON, _ := json.Marshal(talents)

	abilityArray := garray.NewArray()
	ability := gmap.New()
	//技能
	doc.Find("#accordion>div").Each(func(i int, s *goquery.Selection) {
		if i > 3 {
			if i%2 == 0 {
				ability.Set("name", strings.TrimSpace(s.Text()))
				src, _ := s.Find("img").First().Attr("src")
				ability.Set("img", src)
			} else {
				abilitieInfos := garray.NewArray(true)
				//ability desc
				ability.Set("desc", s.Children().First().Text())
				//ability info
				abilitieInfos.Append(map[string]string{
					"label": "mana",
					"value": s.Find(".mana").First().Text(),
				})

				abilitieInfos.Append(map[string]string{
					"label": "cooldown",
					"value": s.Find(".cooldown").First().Text(),
				})

				html, _ := s.Html()
				reg := regexp.MustCompile(`<br/>(?s:(.*?))<span class="attribVal">(?s:(.*?))</span>`)
				if reg != nil {
					result := reg.FindAllStringSubmatch(html, -1)
					for _, text := range result {
						reg, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
						text[2] = reg.ReplaceAllString(text[2], "")
						abilitieInfos.Append(map[string]string{
							"label": strings.TrimSpace(text[1]),
							"value": strings.TrimSpace(text[2]),
						})
					}
					ability.Set("abilitieInfos", abilitieInfos)
					abilityArray.Append(ability)
				}
				ability = gmap.New()
			}
		}
	})
	abilityJSON, _ := json.Marshal(abilityArray)
	//g.Dump(string(abilityJSON))
	return string(talentsJSON), string(abilityJSON)

}
