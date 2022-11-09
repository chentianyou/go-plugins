package main

import (
	"go-plugins/plugins"
	"go-plugins/route"
	"log"
	"net/http"
	"os"
	"path"
	"plugin"
	"strings"
	"sync"
)

func main() {
	pluginHome := "./plugins"
	infos, err := os.ReadDir(pluginHome)
	if err != nil {
		log.Fatalln(err)
	}
	routes := route.Routes{}
	for _, info := range infos {
		if strings.HasSuffix(info.Name(), ".so") {
			pluginPath := path.Join(pluginHome, info.Name())
			pluginHandler, err := plugin.Open(pluginPath)
			if err != nil {
				log.Printf("Warning: %v\n", err)
				continue
			}
			pluginSym, err := pluginHandler.Lookup("PluginImpl")
			if err != nil {
				log.Printf("Warning: %v\n", err)
				continue
			}
			pluginImpl := pluginSym.(plugins.Plugin)
			log.Printf("Info: load plugin [%s] from '%s'\n", pluginImpl.Name(), pluginPath)
			routes = append(routes, pluginImpl.Routes()...)
		}
	}
	RESTRouter := route.NewRouter(routes, "/api/")
	RESTRouter.Handle("/apis", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var apis []string
			for _, r := range routes {
				apis = append(apis, "/api/"+r.Pattern)
			}
			w.Write([]byte(strings.Join(apis, "\n") + "\n"))
		}
	}))
	httpServer := &http.Server{
		Addr:    ":1888",
		Handler: RESTRouter,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		log.Printf("Go plugin server started. Listen at %s.\n", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Could not listen on %s: %v\n", httpServer.Addr, err)
		} else {
			log.Println("Server stopped")
		}
	}()
	wg.Wait()
}
