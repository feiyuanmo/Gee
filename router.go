package gee

import (
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	log.InfofB("Route %4s - %s", method, path)
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	if c.Path == "/favicon.ico" {

	} else {
		key := c.Method + "-" + c.Path
		log.InfofW("IP:%s Method:%s Path:%s", c.Host, c.Method, c.Path)
		if hadler, ok := r.handlers[key]; ok {
			hadler(c)
		} else {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		}
	}
}
