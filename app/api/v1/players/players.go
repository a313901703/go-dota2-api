package players

import (
	"dotaapi/app/service/players"
	"dotaapi/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

//Controller 管理对象
type Controller struct{}

// View 玩家信息
// @produce json
func (c *Controller) View(r *ghttp.Request) {
	playerID := r.GetRouterString("playerID")
	if playerID == "" {
		response.JsonExit(r, 1, "无效的玩家信息")
	}
	var expects []string
	data := players.GetInfo(playerID, expects)
	response.JsonExit(r, 0, "ok", data)
}

// Matches 玩家比赛信息
func (c *Controller) Matches(r *ghttp.Request) {
	playerID := r.GetRouterString("playerID")
	page := r.Get("page", "1")
	pageSize := r.Get("pageSize", "10")
	data, _ := players.Matches(playerID, gconv.Int(pageSize), gconv.Int(page))
	response.JsonExit(r, 0, "ok", data)
}
