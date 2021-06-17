package httprequest

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

var funDataURL string = "http://api.varena.com/"

//FuncDataRequest  fundata api request
func FuncDataRequest(config Config) (ret interface{}) {
	uri := strings.Trim(config.URL, "/")
	url := funDataURL + uri
	config.headers = generateFunHeader(uri, config.Data)
	if config.Data != nil && len(config.Data) > 0 {
		url += "?"
		for k, v := range config.Data {
			url += k + "=" + v.(string)
		}
	}
	response := Request(url, config)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(response), &data)
	if err != nil || response == "" {
		g.Log().Error(fmt.Sprintf("FuncDataRequest JsonToMap err: %s \n", err))
		return
	}
	var retcode = data["retcode"].(float64)
	if retcode != 200 {
		g.Log().Error("FuncDataRequest err ,response code = " + strconv.FormatFloat(retcode, 'E', -1, 64))
		return
	}
	ret = data["data"]
	return
}

//生成请求header请求头
func generateFunHeader(uri string, params map[string]interface{}) map[string]string {
	rand := rand.Int()
	time := time.Now().Unix()
	secret := g.Cfg().Get("dota.secret")
	uri = "/" + uri
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	paramsStr := ""
	for _, k := range keys {
		paramsStr += ("&" + k + "=" + params[k].(string))
	}
	paramsStr = strings.Trim(paramsStr, "&")
	sign := fmt.Sprintf("%s|%s|%s|%s|%s", strconv.Itoa(rand), secret, strconv.FormatInt(time, 10), uri, paramsStr)
	// fmt.Println("sign = ", sign)
	sign = fmt.Sprintf("%x", md5.Sum([]byte(sign)))
	// build header
	headers := map[string]string{
		"Content-Type":    "application/json; charset=utf-8",
		"Accept-ApiKey":   g.Cfg().Get("dota.key").(string),
		"Accept-ApiNonce": strconv.Itoa(rand),
		"Accept-ApiTime":  strconv.FormatInt(time, 10),
		"Accept-ApiSign":  sign,
	}
	return headers
}
