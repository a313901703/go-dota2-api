package news

import (
	baseService "dotaapi/app/service/news"
	"dotaapi/library/response"

	"dotaapi/app/api/v1/basecontroller"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

//Controller 管理对象
type Controller struct {
	basecontroller.Controller
}

//Index 新闻列表
// @produce json
func (c *Controller) Index(r *ghttp.Request) {
	page := r.Get("page", "1")
	pageSize := r.Get("pageSize", "10")
	var service baseService.Service
	service.Page = gconv.Int(page)
	service.PageSize = gconv.Int(pageSize)
	data := service.GetNews()
	response.JsonExit(r, 0, "ok", data)
}
