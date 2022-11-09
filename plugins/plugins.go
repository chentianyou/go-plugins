package plugins

import "go-plugins/route"

type Plugin interface {
	Name() string
	Routes() route.Routes
}
