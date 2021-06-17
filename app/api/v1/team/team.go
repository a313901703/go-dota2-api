package team

import (
	"dotaapi/app/api/v1/basecontroller"
	"dotaapi/app/service/team"
	"dotaapi/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

//Controller 实例
type Controller struct {
	basecontroller.Controller
}

//Index teams list
func (c *Controller) Index(r *ghttp.Request) {
	page := r.Get("page", "1")
	pageSize := r.Get("pageSize", "10")
	var service team.Service
	service.Page = gconv.Int(page)
	service.PageSize = gconv.Int(pageSize)
	data := service.GetAll()
	response.JsonExit(r, 0, "ok", data)
}

//View teamInfo
func (c *Controller) View(r *ghttp.Request) {
	ID := r.GetRouterString("ID")
	if ID == "" {
		response.JsonExit(r, 1, "无效的team ID")
	}
	var service team.Service
	data, err := service.GetInfo(ID)
	if err != nil {
		response.JsonExit(r, 1, "无效的team ID")
	}
	response.JsonExit(r, 0, "ok", data)
}
