package match

import (
	"dotaapi/app/service/match"
	"dotaapi/library/response"

	// "github.com/gogf/gf/net/ghttp"
	"dotaapi/app/api/v1/basecontroller"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 管理对象
type Controller struct {
	basecontroller.Controller
}

//View 比赛详情
// @produce json
func (c *Controller) View(r *ghttp.Request) {
	matchID := r.GetRouterString("matchID")
	if matchID == "" {
		response.JsonExit(r, 1, "无效的比赛")
	}
	service := match.Service{}
	data := service.FindOne(matchID)
	response.JsonExit(r, 0, "ok", data)
}
