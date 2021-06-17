package spider

import (
	"dotaapi/app/model/item"
	"dotaapi/library/httprequest"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/container/garray"

	//"encoding/json"
	"github.com/gogf/gf/frame/g"
)

//GetItems 物品信息
func GetItems() {
	// data, _ := g.Redis().DoVar("GET", keyItemsInfo)
	// if data.String() != "" {
	// 	return
	// }
	config := httprequest.Config{URL: itemURL}
	resp := httprequest.FuncDataRequest(config)
	if resp == nil {
		g.Log().Error("sync items error")
		return
	}
	list := resp.([]interface{})
	//var itemsMap = make([]*item.InfoEntity, len(list))
	itemsMap := make(map[string]*item.InfoEntity)
	for _, v := range list {
		entity := item.NewEntity(v.(map[string]interface{}))
		itemInfoEntity := &item.InfoEntity{Entity: entity}
		itemsMap[entity.Name] = itemInfoEntity
	}
	//物品列表
	spider := Spider{URL: itemsURL}
	doc := spider.GetDoc()
	//WaitGroup
	var wg sync.WaitGroup

	doc.Find(".table-list tbody tr").Each(func(i int, s *goquery.Selection) {
		children := s.Children()
		td1 := children.First()
		td2 := children.Next()
		td3 := children.Last()

		src, _ := td1.Find(".hero-img-list").First().Attr("src")
		name := strings.TrimPrefix(src, "http://cdn.dota2.com/apps/dota2/images/items/")
		name = strings.TrimSuffix(name, "_lg.png")
		if itemsMap[name] == nil {
			fmt.Println(name)
		} else {
			wg.Add(1)
			//记录图片   场次   胜率
			itemsMap[name].Thumb = src
			itemsMap[name].Counts = td2.Find("div").First().Text()
			itemsMap[name].WinRate = td3.Find("div").First().Text()
			go func(name string, ID float64) {
				html, parts := parseItemInfo(name, ID)
				itemsMap[name].Desc = html
				itemsMap[name].Elements = parts
				wg.Done()
			}(name, itemsMap[name].ID)
		}
	})
	wg.Wait()
	itemArr := garray.NewArray()
	for _, v := range itemsMap {
		itemArr.Append(v)
	}
	jsonData, _ := json.Marshal(itemArr)
	g.Redis().DoVar("SET", keyItemsInfo, string(jsonData))
	fmt.Println("sync items info down")
}

func parseItemInfo(name string, ID float64) (string, string) {
	spider := Spider{URL: "http://dotamax.com/item/detail/" + name}
	doc := spider.GetDoc()
	html := ""
	parts := garray.NewStrArray()
	doc.Find(".new-box .iconTooltip").Each(func(i int, s *goquery.Selection) {
		html, _ = s.Html()
		//html = strings.Replace(html, "http", "https", -1)
		s.Find(".iconTooltip_constr .img-shadow").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			parts.Append(src)
		})
	})
	partsJSON, _ := json.Marshal(parts)
	return html, string(partsJSON)
}
