package main

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
	"runtime"
	_ "tech/model"
	"tech/modules/setting"
	"tech/router"
	"tech/modules/middleware"
)

const APP_VER = "2017.10.10"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public", macaron.StaticOptions{
		SkipLogging: true,
	}))
	m.Use(macaron.Renderer())
	m.Use(func(ctx *macaron.Context) {
		ctx.Resp.Header().Set("Access-Control-Allow-Origin","*")
		ctx.Resp.Header().Set("Access-Control-Allow-Methods","POST,GET,OPTIONS,DELETE")
		ctx.Resp.Header().Set("Access-Control-Allow-Headers","x-requested-with,content-type")
	})
	m.Use(middleware.Contexter())
	m.Options("*", func(ctx *macaron.Context) {
	})

	m.Get("/favicon.ico", func(ctx *macaron.Context) {
		ctx.Redirect("/img/favicon.png")
	})
	// Routers
	m.Group("/api", func() {
		m.Group("/v1", func() {
			m.Post("/all", router.All)
			m.Post("/detail", router.Detail)
		})
	},router.IdFilter)
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.HttpPort)
	if err := http.ListenAndServe(listenAddr, m); err != nil {
		log.Fatal("fail to start server")
	}
}
