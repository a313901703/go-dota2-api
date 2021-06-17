package main

import (
	_ "dotaapi/boot"
	_ "dotaapi/router"
	_ "dotaapi/spider"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
