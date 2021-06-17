package router

import (
	"dotaapi/app/api/v1/hero"
	"dotaapi/app/api/v1/item"
	"dotaapi/app/api/v1/league"
	"dotaapi/app/api/v1/match"
	"dotaapi/app/api/v1/news"
	"dotaapi/app/api/v1/players"
	"dotaapi/app/api/v1/team"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// 分组路由注册方式
	s.Group("/v1", func(group *ghttp.RouterGroup) {
		ctHeroes := new(hero.Controller)
		//group.ALL("/hero", new(hero.Controller))
		group.Group("/hero", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctHeroes, "Index")
			group.ALL("/:heroID", ctHeroes, "View")
		})
		//teams
		ctTeams := new(team.Controller)
		group.Group("/teams", func(group *ghttp.RouterGroup) {
			group.ALL("/", ctTeams, "Index")
			group.ALL("/:ID", ctTeams, "View")
		})
		group.ALL("/leagues", new(league.Controller))
		group.ALL("/news", new(news.Controller))
		group.ALL("/item", new(item.Controller))
		group.ALL("/players/:playerID", new(players.Controller), "View")
		group.ALL("/players/:playerID/matches", new(players.Controller), "Matches")
		group.ALL("/match/:matchID", new(match.Controller), "View")
	})
}
