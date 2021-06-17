package httprequest

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	timeOutS = 15
)

//Config http request config
type Config struct {
	URL     string
	Data    map[string]interface{}
	headers map[string]string
}

//Request http request
func Request(url string, config Config) string {
	// request start
	g.Log().Info("request start：" + url)
	// g.Log().Info(config)

	client := ghttp.NewClient()
	client.SetHeader("Content-Type", "application/json; charset=utf-8")
	if len(config.headers) > 0 {
		client.SetHeaderMap(config.headers)
	}
	timeout := time.Duration(timeOutS) * time.Second
	resp := client.SetTimeout(timeout).GetContent(url)
	//fmt.Println(resp)
	// g.Log().Info("request end,response data:" + resp)
	return resp
}
