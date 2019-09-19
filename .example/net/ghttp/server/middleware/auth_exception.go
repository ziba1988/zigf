package main

import (
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	token := r.Get("token")
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func main() {
	s := g.Server()
	s.Group("/admin", func(g *ghttp.RouterGroup) {
		g.MiddlewarePattern("/*action", func(r *ghttp.Request) {
			if action := r.GetRouterString("action"); action != "" {
				switch action {
				case "login":
					r.Middleware.Next()
					return
				}
			}
			MiddlewareAuth(r)
		})
		g.ALL("/login", func(r *ghttp.Request) {
			r.Response.Write("login")
		})
		g.ALL("/dashboard", func(r *ghttp.Request) {
			r.Response.Write("dashboard")
		})
		g.ALL("/user", func(r *ghttp.Request) {
			r.Response.Write("user")
		})
	})
	s.SetPort(8199)
	s.Run()
}