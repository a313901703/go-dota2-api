package spider

import (
	newsModel "dotaapi/app/model/news"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/util/gconv"
)

//GetNews 新闻数据
func GetNews() {
	// data, _ := g.Redis().DoVar("GET", keyNews)
	// if data.String() != "" {
	// 	return
	// }
	entities := garray.NewArray(true)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			spider := Spider{URL: fmt.Sprintf(newsURL, gconv.String(i))}
			doc := spider.GetDoc()
			doc.Find(".news .news_lists .panes .active a").Each(func(i int, s *goquery.Selection) {
				//var entity *newsModel.Entity
				entity := new(newsModel.Entity)
				src, _ := s.Find(".news_logo img").First().Attr("src")
				title := s.Find(".news_msg .title").First().Text()
				content := s.Find(".content").First().Text()
				date := s.Find(".news_msg .date").First().Text()
				entity.Title = title
				entity.Content = content
				entity.Date = date
				entity.Thumb = src
				entities.Append(entity)
				//g.Dump(entity)
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
	if entities.Len() > 0 {
		b, _ := json.Marshal(entities)
		g.Redis().DoVar("SET", keyNews, string(b))
		fmt.Println("sync news success")
	} else {
		fmt.Println("sync news fail")
	}

}
