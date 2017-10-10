package main

import (
	"runtime"
	"tech/modules/setting"
	"gopkg.in/macaron.v1"
	"tech/router"
	"fmt"
	"net/http"
	"log"
	_ "tech/model"
	"tech/seed"
)

const APP_VER = "2017.10.10"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}


func main () {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public",macaron.StaticOptions{
		SkipLogging:true,
	}))

	m.Get("/favicon.ico", func(ctx *macaron.Context) {
		ctx.Redirect("/img/favicon.png")
	})

	// Routers
	m.Group("/api", func() {
		m.Group("/v1", func() {
			m.Post("/all",router.All)
		})
	})
	listenAddr := fmt.Sprintf("0.0.0.0:%d" ,setting.HttpPort)
	if err := http.ListenAndServe(listenAddr ,m) ; err != nil {
		log.Fatal("fail to start server")
	}
}


















