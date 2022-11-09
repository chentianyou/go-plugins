package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
	ACCheck     bool // 是否做权限认证
	Audit       bool
	NamePath    string
	ACRules     []string
}

type Routes []Route

func NewRouter(routes Routes, prefix string) *http.ServeMux {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method...).
			Path(prefix + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	serverMux := http.NewServeMux()
	serverMux.Handle("/", router)

	return serverMux
}