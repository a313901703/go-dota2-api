package httprequest

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gogf/gf/frame/g"
)

var openDotaURL string = "https://api.opendota.com/api/"

//OpenDotaRequest  open dota api request
func OpenDotaRequest(config Config) interface{} {
	url := openDotaURL + strings.Trim(config.URL, "/")
	response := Request(url, config)
	var data interface{}
	if response != "" {
		err := json.Unmarshal([]byte(response), &data)
		if err != nil {
			g.Log().Info(fmt.Sprintf("OpenDotaRequest JsonToMap err: %s \n", err))
		}
	}

	return data
}

//OpenDotaRequestAsStr  open dota api request
func OpenDotaRequestAsStr(config Config) string {
	url := openDotaURL + strings.Trim(config.URL, "/")
	response := Request(url, config)

	return response
}
