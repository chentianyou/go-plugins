package main

import (
	"go-plugins/route"
	"net/http"
)

var PluginImpl = Module1Plugin{}

type Module1Plugin struct {
	routes route.Routes
}

func (h *Module1Plugin) Name() string {
	return "module1_plugin"
}

func (h *Module1Plugin) Routes() route.Routes {
	return h.routes
}

func (h *Module1Plugin) registerRoute(r route.Route) {
	h.routes = append(h.routes, r)
}

func ModulePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("欢迎来到Module1~\n"))
	}
}

func init() {
	PluginImpl.registerRoute(route.Route{
		Name: "ModulePage",
		Method:      []string{http.MethodGet},
		Pattern:     "module1",
		HandlerFunc: ModulePage,
	})
}