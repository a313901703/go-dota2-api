package hero

import (
	"dotaapi/app/service/heroes"
	"dotaapi/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Controller heroes管理对象
type Controller struct{}

// Index 英雄列表数据
// @success 200 {object} response.JsonResponse "执行结果"
func (c *Controller) Index(r *ghttp.Request) {
	resp := heroes.GetHeroes()
	response.JsonExit(r, 0, "ok", resp)
}

// View 英雄详情
// @success 200 {object} response.JsonResponse "执行结果"
func (c *Controller) View(r *ghttp.Request) {
	heroID := r.GetRouterString("heroID")
	if heroID == "" {
		response.JsonExit(r, 1, "无效的heroID1")
	}
	resp, isOk := heroes.GetHero(heroID)
	if !isOk {
		response.JsonExit(r, 1, "无效的heroID2")
	}
	response.JsonExit(r, 0, "ok", resp)
}

// Ability 技能
func (c *Controller) Ability(r *ghttp.Request) {
	resp, _ := heroes.GetHero("1")
	response.JsonExit(r, 0, "ok", resp)
}
