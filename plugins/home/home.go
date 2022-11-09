package main

import (
	"go-plugins/route"
	"net/http"
)

var PluginImpl = HomePlugin{}

type HomePlugin struct {
	routes route.Routes
}

func (h *HomePlugin) Name() string {
	return "home_plugin"
}

func (h *HomePlugin) Routes() route.Routes {
	return h.routes
}

func (h *HomePlugin) registerRoute(r route.Route) {
	h.routes = append(h.routes, r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("欢迎来到首页~\n"))
	}
}

func init() {
	PluginImpl.registerRoute(route.Route{
		Name: "HomePage",
		Method:      []string{http.MethodGet},
		Pattern:     "home",
		HandlerFunc: HomePage,
	})
}