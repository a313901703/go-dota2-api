package item

import (
	"dotaapi/app/service/items"
	"dotaapi/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//Controller 管理对象
type Controller struct{}

// Index 物品列表
// @produce json
// @success 200 {object} response.JsonResponse "执行结果"
func (c *Controller) Index(r *ghttp.Request) {
	resp := items.Fetch()
	response.JsonExit(r, 0, "ok", resp)
}
